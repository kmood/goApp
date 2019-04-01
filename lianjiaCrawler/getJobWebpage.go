package job_crawler_project

type requestFake struct {
	method  string
	url     string
	headers map[string]string
}
