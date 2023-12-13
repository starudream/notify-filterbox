package filterbox

type Message struct {
	Title   string `json:"android.title"`
	Text    string `json:"android.text"`
	BigText string `json:"android.bigText"`

	AppName     string `json:"filterbox.field.APP_NAME"`
	PackageName string `json:"filterbox.field.PACKAGE_NAME"`
	ChannelId   string `json:"filterbox.field.CHANNEL_ID"`
	Ongoing     bool   `json:"filterbox.field.ONGOING"`
	When        int    `json:"filterbox.field.WHEN"`
}
