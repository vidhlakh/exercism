package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
	"text/tabwriter"
)

type Team struct {
	name   string
	wins   int
	losses int
	draws  int
}

func (t Team) points() int {
	return t.wins*3 + t.draws
}

func (t Team) mp() int {
	return t.wins + t.draws + t.losses
}

type Tournament []Team

func (t Tournament) Len() int      { return len(t) }
func (t Tournament) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t Tournament) Less(i, j int) bool {
	if t[i].points() == t[j].points() {
		if t[i].name > t[j].name {
			return true
		}
		return false
	}

	return t[i].points() < t[j].points()
}

func Tally(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	teams := make(map[string]*Team)

	for scanner.Scan() {
		st := scanner.Text()
		if st == "" || strings.HasPrefix(st, "#") {
			continue
		}

		itens := strings.Split(st, ";")
		if len(itens) != 3 {
			return fmt.Errorf("invalid input format")
		}

		home, away, result := itens[0], itens[1], itens[2]
		checkIfExists(home, teams)
		checkIfExists(away, teams)

		switch result {
		case "win":
			teams[home].wins++
			teams[away].losses++
		case "loss":
			teams[away].wins++
			teams[home].losses++
		case "draw":
			teams[home].draws++
			teams[away].draws++
		default:
			return fmt.Errorf("invalid match result")
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	printResults(w, rank(teams))
	return nil
}

func checkIfExists(team string, teams map[string]*Team) {
	if _, exists := teams[team]; !exists {
		t := Team{}
		t.name = team
		teams[team] = &t
	}
}

func rank(teams map[string]*Team) Tournament {
	ts := make([]Team, 0, len(teams))

	for _, team := range teams {
		ts = append(ts, *team)
	}

	tournament := Tournament(ts)
	sort.Sort(sort.Reverse(tournament))

	return tournament
}

func printResults(w io.Writer, tournament Tournament) {
	tab := tabwriter.NewWriter(w, 1, 0, 8, ' ', 0)
	defer tab.Flush()

	fmt.Fprintln(tab, "Team\t| MP |  W |  D |  L |  P")
	for _, team := range tournament {
		fmt.Fprintf(tab, "%s\t| %2d | %2d | %2d | %2d | %2d\n", team.name, team.mp(), team.wins, team.draws, team.losses, team.points())
	}
}
