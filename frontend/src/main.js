import Vue from 'vue'
import App from './App.vue'
import { BootstrapVue} from 'bootstrap-vue'
//import axios from 'axios';

// Import Bootstrap an BootstrapVue CSS files (order is important)
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.use(BootstrapVue)

Vue.config.productionTip = false
Vue.use(BootstrapVue)

/*const base = axios.create({
  baseURL: "http://localhost:8080/api"
})*/

new Vue({
  render: h => h(App),
  components:{
    App
  }
}).$mount('#app')
