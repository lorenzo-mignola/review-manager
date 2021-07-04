import { createApp } from 'vue'
import WaveUI from 'wave-ui'
import 'wave-ui/dist/wave-ui.css'
import App from './App.vue'
import router from './router'
import store, { key } from './store'

const app = createApp(App)
new WaveUI(app, {
  colors: {
    primary: '#8b9716',
    secondary: '#673ab7',
  },
})
app.use(store, key)
app.use(router)
app.mount('#app')
