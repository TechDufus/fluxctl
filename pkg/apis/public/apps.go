package apis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type appDataPorts struct {
	IP          string `json:"IP,omitempty"`
	PrivatePort int    `json:"PrivatePort,omitempty"`
	PublicPort  int    `json:"PublicPort,omitempty"`
	Type        string `json:"Type"`
}

type appData struct {
	ID      string            `json:"Id"`
	Names   []string          `json:"Names"`
	Image   string            `json:"Image"`
	ImageID string            `json:"ImageID"`
	Command string            `json:"Command"`
	Created int               `json:"Created"`
	Ports   []appDataPorts    `json:"Ports"`
	Labels  map[string]string `json:"Labels"`
	State   string            `json:"State"`
	Status  string            `json:"Status"`
}

type ListAllApps struct {
	Status string    `json:"status"`
	Data   []appData `json:"data"`
}

type PublicAPI struct {
	Apps PublicAPIApps
}
type PublicAPIApps struct {
	ListAllApps string
}

var PublicAPIs PublicAPI = PublicAPI{
	Apps: PublicAPIApps{
		ListAllApps: "https://api.runonflux.io/apps/listallapps",
	},
}

func GetListAllApps() (ListAllApps, error) {
	resp, _ := http.Get(PublicAPIs.Apps.ListAllApps)
	body, _ := io.ReadAll(resp.Body)
	var data ListAllApps
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return ListAllApps{}, err
	}

	return data, nil
}

func FormatListAllApps(data []appData, format string) {
	switch format {
	case "default":
		formatListAllAppsDefault(data)
	case "wide":
		formatListAllAppsWide(data)
	case "json":
		formatListAllAppsJSON(data)
	default:
		formatListAllAppsDefault(data)
	}
}

func formatListAllAppsJSON(data []appData) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}
	fmt.Println(string(jsonData))
}

func formatListAllAppsWide(data []appData) {
	header := "ID\tNames\tImage\tCommand\tCreated\tPorts\tLabels\tState\tStatus\n"
	fmt.Printf(header)
	for _, app := range data {
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", app.ID, app.Names, app.Image, app.Command, app.Created, app.Ports, app.Labels, app.State, app.Status)
	}
}

func formatListAllAppsDefault(data []appData) {
	header := "Names\tImage\tState\tStatus\n"
	fmt.Printf(header)
	for _, app := range data {
		fmt.Printf("%s\t%s\t%s\t%s\n", app.Names[0], app.Image, app.State, app.Status)
	}
}
