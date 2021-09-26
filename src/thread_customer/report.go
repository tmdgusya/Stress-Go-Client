package threadcustomer

type Report struct {
	fail_count int
	success_count int
}

func (r *Report) SuccessTestResult() {
	r.success_count += 1
}

func (r *Report) FailTestResult() {
	r.fail_count += 1
}

func (r *Report) GetSuccessCount() int {
	return r.success_count
}

func (r *Report) GetFailCount() int {
	return r.fail_count
}

func ReportFactory() *Report {
	return &Report {fail_count: 0, success_count: 0};
}