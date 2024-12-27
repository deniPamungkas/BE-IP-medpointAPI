# BE Medpoint System
repository ini adalah repository setup proJect medpoint sistem untuk bagian backend api application. Ini adalah setup project yang menggunakan teknology golang dan raiden framework serta supabase sebagai database as a serveice. untuk database, saya telah mengkoneksikannya dengan supabase cloud.

## teknologi
- Language -> Golang
- Framework -> Raiden
- Database -> Supabase cloud

## struktur folder
- CMS-IP-medpoint/
- ├── configs
- │   ├── app.yaml          # App configuration file
- │   └── modules/          # Module configuration file
- │       ├── module_a.yaml
- │       ├── module_b.yaml
- │       └── ...
- ├── cmd
- │   └── project-name
- │       └── project_name.go    # Main project
- │   └── apply/main.go
- │   └── import/main.go
- ├── internal
- │   ├── bootstrap
- │   │   ├── route.go
- │   │   ├── rpc.go
- │   │   ├── roles.go
- │   │   ├── models.go
- │   │   └── storages.go
- │   ├── controllers
- │   │   └── hello.go    # Example controller
- │   ├── roles           # ACL/RLS definition
- │   │   ├── members.go
- │   │   └── ...
- │   ├── models          # All model will auto-sync
- │   │   ├── users.go
- │   │   └── ...
- │   ├── rpc             # RPC
- │   │   └── get_user.go
- │   ├── topics          # Real-time
- │   │   ├── notification.go
- │   │   └── inbox.go
- │   │
- │   └── storages
- │       └── avatar.go
- ├── build
- │   └── state      # Bytecode containing raiden state
- ├── go.sum
- └── go.mod

## How to Run the project in your computer 💻

### Requirements
- install raiden framework -> https://raiden.sev-2.com/docs/installation
- install golang -> https://go.dev/doc/install

### 3. run backend 
```bash
cd server
```
```bash
npm install
```
