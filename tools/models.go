package tools

import "time"

type Html struct {
	HtmlBase
	PersonBase
	Work         []Work    `json:"work"`
	LeftAbility  []Ability `json:"left_ability"`
	RightAbility []Ability `json:"right_ability"`
	Project      []Project `json:"project"`
}

func NewHtml() *Html {
	h := &Html{}
	return h
}

type HtmlBase struct {
	Title string `json:"title"`
	Update
	DownloadPDF string `json:"download_pdf"`
	FooterLink  string `json:"footer_link"`
	FooterName  string `json:"footer_name"`
}

type Update struct {
	UpdateYear  int `json:"update_year"`
	UpdateMonth int `json:"update_month"`
}

func (u *Update) Now() {
	now := time.Now()
	u.UpdateYear = now.Year()
	u.UpdateMonth = int(now.Month())
}

type PersonBase struct {
	Name       string `json:"name"`
	WorkJob    string `json:"work_job"`
	Sex        string `json:"sex"`
	Brithday   string `json:"brithday"`
	School     string `json:"school"`
	Specialty  string `json:"specialty"`
	Degree     string `json:"degree"`
	Finish     string `json:"finish"`
	Experience int    `json:"experience"`
	BlogUrl    string `json:"blog_url"`
	BlogName   string `json:"blog_name"`
	GitAddress string `json:"git_address"`
	GitName    string `json:"git_name"`
	GitAccount string `json:"git_account"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

type Work struct {
	Company string   `json:"company"`
	Start   string   `json:"start"`
	End     string   `json:"end"`
	Url     string   `json:"url"`
	Favicon string   `json:"favicon"`
	At      string   `json:"at"`
	Working []string `json:"working"`
}

type Ability struct {
	AbilityName      string   `json:"ability_name"`
	AbilityRecommend []string `json:"ability_recommend"`
}

type Project struct {
	ProjectType string      `json:"project_type"`
	ProjectUrl  string      `json:"project_url"`
	Icon        string      `json:"icon"`
	Introduce   []Introduce `json:"introduce"`
}

type Introduce struct {
	Addr string `json:"addr"`
	Tag  string `json:"tag"`
	Mark string `json:"mark"`
}
