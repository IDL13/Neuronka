import requests
from bs4 import BeautifulSoup


def request(url):
    praises_x = []
    praises_y = []
    r = requests.get(url).text
    soup1 = BeautifulSoup(r, 'lxml')
    x = soup1.findAll("div", class_ = "js_td_width")

    soup2 = BeautifulSoup(r, 'lxml')
    y = soup2.findAll("div", class_ = "js_t_w")

    for y_iter in reversed(range(1, 4)):
        art_y = y[y_iter].text
        praises_y.append(art_y.split(" ")[0].split(".")[0])

    praises_y = list(map(float, praises_y))

    print(praises_y)

    for i in reversed(range(0, 3)):
        art = x[i].text
        praises_x.append(art.split(" ")[1])   
    
    praises_x = list(map(float, praises_x))

    print(praises_x)

def main():
    request("https://investfunds.ru/indexes/9021/") 

if __name__ == "__main__":
    main()
