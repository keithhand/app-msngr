package application

type application struct {
	helmChart  helmChart
	kubernetes kubernetes
}

type kubernetes interface {
	Logs() error
}

type helmChart interface {
	Install() error
	Update() error
	Uninstall() error
}

func New(helm helmChart, k8s kubernetes) *application {
	return &application{
		helmChart:  helm,
		kubernetes: k8s,
	}
}

func (a *application) Deploy() error {
	return a.helmChart.Install()
}

func (a *application) Update() error {
	return a.helmChart.Update()
}

func (a *application) Logs() error {
	return a.kubernetes.Logs()
}

func (a *application) Reinit() error {
	if err := a.helmChart.Uninstall(); err != nil {
		return err
	}

	if err := a.helmChart.Install(); err != nil {
		return err
	}
	return nil
}

func (a *application) Destroy() error {
	return a.helmChart.Uninstall()
}
