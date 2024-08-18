<template>
  <div class="homepage">


    <section class="hero">
      <div class="hero-content">
        <h1>Welcome to Cinema</h1>
        <p>Watch your favorite movies and shows in high quality</p>
        <button class="cta-button">Explore Now</button>
      </div>
    </section>

    <section class="movies-section">
      <h2>Featured Movies</h2>
      <div class="movies-grid">
        <div class="movie-card" v-for="movie in movies" :key="movie.id">
          <img :src="movie.image" :alt="movie.title" />
          <div class="movie-info">
            <h3>{{ movie.title }}</h3>
            <p>{{ movie.genre }}</p>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
<style scoped>
.homepage {
  font-family: 'Arial', sans-serif;
  color: #fff;
  background: linear-gradient(135deg, #0f0f0f, #1a1a1a);
  padding: 0;
  margin: 0;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  background-color: #181818;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.5);
}

.logo {
  font-size: 1.8em;
  font-weight: bold;
  color: #fff;
}

.navbar ul {
  list-style: none;
  display: flex;
  gap: 20px;
  padding: 0;
  margin: 0;
}

.navbar ul li a {
  text-decoration: none;
  color: #aaa;
  transition: color 0.3s;
}

.navbar ul li a:hover {
  color: #fff;
}

.hero {
  height: 80vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: url('your-hero-background-image.jpg') no-repeat center center/cover;
  position: relative;
  text-align: center;
}

.hero::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
}

.hero-content {
  z-index: 1;
}

.hero h1 {
  font-size: 3em;
  margin: 0;
}

.hero p {
  font-size: 1.2em;
  margin: 10px 0 20px;
}

.cta-button {
  background-color: #ff4d4d;
  color: #fff;
  padding: 15px 30px;
  border: none;
  border-radius: 30px;
  cursor: pointer;
  font-size: 1em;
  transition: background-color 0.3s, transform 0.3s;
}

.cta-button:hover {
  background-color: #ff3333;
  transform: scale(1.05);
}

.movies-section {
  padding: 50px 20px;
}

.movies-section h2 {
  text-align: center;
  font-size: 2em;
  margin-bottom: 30px;
}

.movies-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
}

.movie-card {
  background-color: #222;
  border-radius: 10px;
  overflow: hidden;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.3);
  transition: transform 0.3s;
}

.movie-card img {
  width: 100%;
  height: auto;
  display: block;
}

.movie-info {
  padding: 15px;
}

.movie-info h3 {
  margin: 0;
  font-size: 1.2em;
}

.movie-info p {
  margin: 5px 0 0;
  font-size: 0.9em;
  color: #bbb;
}

.movie-card:hover {
  transform: translateY(-10px);
}
</style>
