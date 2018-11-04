var app = new Vue({
  el: "#app",
  data: {
      questions:[],
      didWell: false,
      successRate: 0,
      correctAnswers: 0,
      submited: false
  },
  methods: {
      getQuestions() {
        var that = this;
        axios.get('http://localhost:8080/quiz')
          .then(function (response) {
            response.data.forEach((q, idx) => {
              Vue.set(that.questions, idx, q)
            });
          })
          .catch(error => console.log(error))
    },
    resetSubmit(){
      this.submited = false
    },
    submit(){
      this.resetSubmit()
      var choices = []
      this.questions.map(q => choices.push(q.choice))
      var that = this
      axios.post('http://localhost:8080/quiz/results', {choices: choices})
        .then(function (response) {
          if (response.status == 201){
            that.submited = true
            that.successRate = response.data.successRate
            that.correctAnswers = response.data.correct
            that.didWell = that.successRate >= 50 ? true : false
          }
        })
        .catch(error => console.log(error))     
    }
  },
  computed:{
    quiz(){
      return this.questions.length > 0 ? this.questions : this.getQuestions()
    }
  }
})