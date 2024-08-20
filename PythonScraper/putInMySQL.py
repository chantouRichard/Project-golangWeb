import json
import pymysql

# 数据库连接配置
db_config = {
    'host': 'localhost',  # 数据库主机地址
    'user': 'root',  # 数据库用户名
    'password': '123456',  # 数据库密码
    'database': 'theater',  # 数据库名
    'charset': 'utf8mb4'  # 字符集
}

def load_json(filename):
    with open(filename, 'r', encoding='utf-8') as file:
        return json.load(file)

def insert_movies_to_db(movies):
    connection = pymysql.connect(**db_config)
    try:
        with connection.cursor() as cursor:
            sql = """
                INSERT INTO movies (title, quote, thumbnail_url, rating)
                VALUES (%s, %s, %s, %s)
            """
            for movie in movies:
                cursor.execute(sql, (movie['title'], movie['quote'], movie['image_url'], movie['rating']))
        connection.commit()
    finally:
        connection.close()

def main():
    filename = 'douban_top250_movies.json'
    movies = load_json(filename)
    insert_movies_to_db(movies)
    print('Movies have been imported into the database.')

if __name__ == '__main__':
    main()
