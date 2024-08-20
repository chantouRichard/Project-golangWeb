# import numpy as np
#
# x = np.float32(1.0)
# print("初始值:", x)
#
# for i in range(149):
#     # x -= np.float32(1.0e-10)
#     x /= 2
# print("减少后的值:", x)
#
# x /= 2
# print(x)

import requests
from bs4 import BeautifulSoup
import json


def get_movie_details(url):
    headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36'
    }
    response = requests.get(url, headers=headers)
    response.encoding = 'utf-8'
    return response.text


def parse_movie_data(html):
    soup = BeautifulSoup(html, 'html.parser')
    movies = []

    for item in soup.find_all('div', class_='item'):
        title = item.find('span', class_='title').get_text()
        image_url = item.find('img')['src']
        rating = item.find('span', class_='rating_num').get_text()
        quote = item.find('span', class_='inq')
        quote_text = quote.get_text() if quote else "No quote"

        # 获取电影分类信息
        genre_tag = item.find('span', class_='genre')
        genre = genre_tag.get_text() if genre_tag else "Unknown"

        movies.append({
            'title': title,
            'image_url': image_url,
            'rating': rating,
            'quote': quote_text,
            'genre': genre
        })

    return movies


def main():
    base_url = 'https://movie.douban.com/top250'
    all_movies = []

    for start in range(0, 250, 25):  # 豆瓣Top250分为10页，每页25条数据
        url = f'{base_url}?start={start}&filter='
        html = get_movie_details(url)
        movies = parse_movie_data(html)
        all_movies.extend(movies)

    with open('douban_top250_movies_with_genres.json', 'w', encoding='utf-8') as f:
        json.dump(all_movies, f, ensure_ascii=False, indent=4)

    print('Movies with genres have been written to douban_top250_movies_with_genres.json')


if __name__ == '__main__':
    main()
