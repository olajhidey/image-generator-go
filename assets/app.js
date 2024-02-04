const { createApp, ref } = Vue


const app = createApp({
    data() {
        return {
            results: [],
            prompt: '',
            size: '',
            variation: 0,
            isLoading: false
        }
    },

    methods: {
        async getImage(){
            try {
                this.isLoading = true

                let request = await axios.post("/api/getImage", {
                    prompt: this.prompt,
                    size: this.size,
                    n: this.variation,
                    model: "dall-e-3"
                })

                if(request){
                    this.isLoading = false
                }

                console.log(request.data)
                this.results = request.data.data.data
                console.log(this.results)
                
            } catch (error) {
                this.isLoading = false
                console.log(error)
            }
        }
    }
})

app.config.compilerOptions.delimiters = ['${', '}']

app.mount('#app')