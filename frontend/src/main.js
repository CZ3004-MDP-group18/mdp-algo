import Vue from 'vue'
import App from './App.vue'
import { BootstrapVue} from 'bootstrap-vue'

// Import Bootstrap an BootstrapVue CSS files (order is important)
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.use(BootstrapVue)

Vue.config.productionTip = false
Vue.use(BootstrapVue)

/*const base = axios.create({
  baseURL: "127.0.0.1:3000/"
})*/

new Vue({
  render: h => h(App),
  components:{
    App
  }
}).$mount('#app')
