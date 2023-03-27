pip install requests
pip install bs4
pip install lxml
cd ..
cd internal/parser
python3 parser.py
cd ..
cd .. 
cp .\internal\parser\*.txt .\cmd
cd cmd
go run main.go