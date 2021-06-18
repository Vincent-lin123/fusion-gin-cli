package new

const TplProjectStructure = `
├── Makefile
├── README.md
├── main.go # 入口文件
├── config_file # 配置文件
│   ├── config.toml
│   ├── menu.yaml
│   └── model.conf
├── docs
│   └── data_model.md
├── go.mod
├── go.sum
├── api
│   └── mock
├── service
├── config
├── contextx
├── ginx
├── middleware
├── model
│   └── gormx
├── module
│   └── adapter
├── router
├── schema
├── swagger
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── test
├── lib
│   ├── auth
│   │   └── jwtauth
│   ├── errors
│   ├── logger
│   │   ├── hook
│   │   │   ├── gorm
│   └── util
│       ├── hash
│       ├── json
│       ├── structure
│       ├── trace
│       ├── uuid
│       └── yaml
└── scripts
`
