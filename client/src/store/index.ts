import { InjectionKey } from 'vue'
import { createStore, Store, useStore as baseUseStore } from 'vuex'
import { Review } from '../types/Review.interface'
import { User } from './../types/Review.interface'

export interface State {
  reviews: Review[]
  currentReview: {
    text: string
    star: number
  }
  currentUser?: User
}

export const key: InjectionKey<Store<State>> = Symbol()

const store = createStore<State>({
  state() {
    return {
      reviews: [],
      currentReview: {
        text: '',
        star: 0,
      },
      currentUser: undefined,
    }
  },
  mutations: {},
})

export function useStore() {
  return baseUseStore(key)
}

export default store
