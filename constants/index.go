package constants

type GitTemplate struct {
	DownloadUrl string
	Description string
	Branch      string
	IsBranch    bool
}

var (
	TemplateArray      = []string{"thin", "no-i18n", "tauri", "admin"}
	TemplateRepoArray  = []string{"github", "gitee"}
	TemplateRepoMapped = map[string][]GitTemplate{
		"thin": {
			{
				DownloadUrl: "https://github.com/xiaoxian521/pure-admin-thin.git", // 模板下载地址
				Description: "vue-pure-admin精简版",                                  // 模板描述
				Branch:      "main",                                               // 分支
				IsBranch:    false,                                                //是否为分支
			},
			{
				DownloadUrl: "https://gitee.com/yiming_chang/pure-admin-thin.git", // 模板下载地址
				Description: "vue-pure-admin精简版",                                  // 模板描述
				Branch:      "main",                                               // 分支
				IsBranch:    false,
			},
		},
		"no-i18n": {
			{
				DownloadUrl: "https://github.com/xiaoxian521/pure-admin-thin.git", // 模板下载地址
				Description: "vue-pure-admin精简版移除国际化",                             // 模板描述
				Branch:      "delete-i18n",                                        // 分支
				IsBranch:    true,
			},
			{
				DownloadUrl: "https://gitee.com/yiming_chang/pure-admin-thin.git", // 模板下载地址
				Description: "vue-pure-admin精简版移除国际化",                             // 模板描述
				Branch:      "delete-i18n",                                        // 分支
				IsBranch:    true,
			},
		},
		"tauri": {
			{
				DownloadUrl: "https://github.com/xiaoxian521/tauri-pure-admin.git", // 模板下载地址
				Description: "vue-pure-admin精简版的tauri模板",                           // 模板描述
				Branch:      "main",                                                // 分支
				IsBranch:    false,
			},
			{
				DownloadUrl: "https://gitee.com/yiming_chang/tauri-pure-admin.git", // 模板下载地址
				Description: "vue-pure-admin精简版的tauri模板",                           // 模板描述
				Branch:      "main",                                                // 分支
				IsBranch:    false,
			},
		},
		"admin": {
			{
				DownloadUrl: "https://github.com/xiaoxian521/vue-pure-admin.git", // 模板下载地址
				Description: "vue-pure-admin精简版",                                 // 模板描述
				Branch:      "main",                                              // 分支
				IsBranch:    false,
			},
			{
				DownloadUrl: "https://gitee.com/yiming_chang/vue-pure-admin.git", // 模板下载地址
				Description: "vue-pure-admin完整版",                                 // 模板描述
				Branch:      "main",                                              // 分支
				IsBranch:    false,
			},
		},
	}
	DefaultVersion   = "last"
	DefaultForce     = false
	DefaultLocalPath = "./"
)
