package servers

import "github.com/Rayato159/neversuitup-e-commerce-test/modules/monitor/monitorHandler"

type IMonitorModule interface {
	Init()
	Handler() monitorHandler.IMonitorHandler
}

type monitorModule struct {
	*module
	handler monitorHandler.IMonitorHandler
}

func (m *module) NewMonitorModule() IMonitorModule {
	return &monitorModule{
		module:  m,
		handler: monitorHandler.NewMonitorHandler(m.s.cfg),
	}
}

func (m *monitorModule) Init() {
	m.s.app.Get("/", m.handler.HealthCheck)
}
func (m *monitorModule) Handler() monitorHandler.IMonitorHandler { return m.handler }
