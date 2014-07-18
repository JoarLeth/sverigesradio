package track

import (
	"code.google.com/p/go.net/html"
	"fmt"
	"github.com/joarleth/sverigesradio/episode"
	"net/http"
	"strings"
)

type Track struct {
	Name    string
	Artists []string
	Album   string
	Uri     string
}

func GetTracks(episode episode.Episode) {
	trackListUrl := fmt.Sprintf("http://sverigesradio.se/sida/latlista.aspx?programid=%d", episode.ExternalProgramId)

	resp, _ := http.Get(trackListUrl)
	defer resp.Body.Close()

	/*if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNotModified) {
		return []Program{}, errors.New(fmt.Sprintf("GET request in GetPrograms returned status %d rather than %d or %d", resp.StatusCode, http.StatusOK, http.StatusNotModified))
	}*/

	doc, _ := html.Parse(resp.Body)
	n := getTracksSection(doc)
	fmt.Println(n)
	n = getTracksTable(n, episode)
	fmt.Println(n)

	// It hangs when I uncomment this
	tracks := getTracks(n)
	_ = tracks
}

func getTracks(n *html.Node) []Track {
	var tracks []Track

	trackNodes := getTrackNodes(n)

	extractTracks(trackNodes)

	return tracks
}

func extractTracks(trackNodes []*html.Node) {
	artistAndNameFinder := func(n *html.Node) *html.Node {
		if n.Type == html.ElementNode && n.Data == "span" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && attr.Val == "track-title" {
					return n
				}
			}
		}
		return nil
	}

	for _, trackNode := range trackNodes {
		n := findFirstMatchingNode(trackNode, artistAndNameFinder, false)

		fmt.Println(n.FirstChild.Data)
	}

	//return tracks
}

func getTrackNodes(n *html.Node) []*html.Node {
	nodeFinder := func(n *html.Node) *html.Node {
		if n.Type == html.ElementNode && n.Data == "li" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && strings.HasPrefix(attr.Val, "track ") {
					return n
				}
			}
		}
		return nil
	}

	var trackNodes []*html.Node

	for {
		n = findFirstMatchingNode(n, nodeFinder, true)
		if n == nil {
			break
		}
		trackNodes = append(trackNodes, n)

		n = n.NextSibling
	}

	return trackNodes
}

func getTracksTable(n *html.Node, episode episode.Episode) *html.Node {
	headerLinkFinder := func(n *html.Node) *html.Node {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" &&
					attr.Val == fmt.Sprintf("/sida/avsnitt/%d?programid=%d",
						episode.ExternalId,
						episode.ExternalProgramId) {
					return n.Parent
				}
			}
		}
		return nil
	}

	trackListFinder := func(n *html.Node) *html.Node {
		if n.Type == html.ElementNode && n.Data == "ul" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && strings.Contains(attr.Val, "track-list") {
					return n
				}
			}
			return n
		}
		return nil
	}

	n = findFirstMatchingNode(n, headerLinkFinder, false)

	return findFirstMatchingNode(n, trackListFinder, true)
}

func getTracksSection(n *html.Node) *html.Node {
	nodeFinder := func(n *html.Node) *html.Node {
		if n.Type == html.ElementNode && n.Data == "ul" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && strings.Contains(attr.Val, "song-list-flow") {
					return n
				}
			}
		}
		return nil
	}

	return findFirstMatchingNode(n, nodeFinder, false)
}

// findFirstMatchingNode walks the node tree recursively and returns the first node returned
// by the nodeFinder function or nil if no match is found
func findFirstMatchingNode(n *html.Node, nodeFinder func(n *html.Node) *html.Node, searchSiblings bool) *html.Node {
	for {
		if n == nil {
			return n
		}
		if t := nodeFinder(n); t != nil {
			return t
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			d := findFirstMatchingNode(c, nodeFinder, searchSiblings)
			if d != nil {
				return d
			}
		}

		if !searchSiblings {
			return nil
		}

		n = n.NextSibling
	}
}
