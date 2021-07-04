export interface Review {
  id: number
  user: User
  stars: number
  text: number
  date: Date
  film: Film
}

export interface User {
  id: number
  username: string
}

export interface Film {
  id: number
  title: string
}
