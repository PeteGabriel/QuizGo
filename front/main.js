var app = new Vue({
  el: "#app",
  data: {
      questions:[
        {
          "key": 1,
          "content":"What is the national animal of China?",
          "options":[
            "Giant Panda",
            "Panda",
            "Red Panda",
            "Polar Bear"
          ],
          choice: "" 
        },
        {
          "key": 2,
          "content":"What is the noisiest city in the world?",
          "options":[
            "Lisbon",
            "London",
            "New York",
            "Hong Kong"
          ],
          choice: ""
        }
      ],
  },
  methods: {
      addToCart() {
          this.cart += 1
      },
     
  },
  computed: {
      title() {
          return this.brand + " " + this.product
      },
  }
})