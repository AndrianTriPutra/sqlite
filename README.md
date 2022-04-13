# Example sqlite

this code is an example using a sqlite database with golang, if you need records while using SBC I recommend this

### instruction
- sudo apt install sqlite3
- git clone https://github.com/AndrianTriPutra/sqlite.git
- cd sqlite
- go run . s_insert
- go run . s_read
- go run . s_update
- go run . s_read
- go run . m_insert
- go run . m_read
- go run . m_delete
- go run . m_csv

### dependency
- gorm.io/gorm
- gorm.io/driver/sqlite
