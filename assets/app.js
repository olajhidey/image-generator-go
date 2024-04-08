const { createApp, ref } = Vue


const app = createApp({
    data() {
        return {
            results: [],
            prompt: '',
            size: '',
            variation: 0,
            model: '',
            style: '',
            isLoading: false
        }
    },

    methods: {
        async getImage(){

            if(this.model == "dall-e-3" && this.variation > 1){
                alert("Only 1 variation is available on the dall-e-3 model")
            }else{
                try {
                    this.isLoading = true
    
                    let request = await axios.post("/api/getImage", {
                        prompt: this.prompt,
                        size: this.size,
                        n: this.variation,
                        style: this.style,
                        model: this.model
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

          
        },

        downloadImage(imageUrl){

            console.log(imageUrl)

            var a = document.createElement("a")
            const filename = new Date().getTime()
            a.href = imageUrl
            // change the filename 
            a.download = `${filename}.jpg`

            // Programmatically initiate the download
            a.click();

            // Remove the anchor element from the body
            document.body.removeChild(a)
        }
    }
})

app.config.compilerOptions.delimiters = ['${', '}']

app.mount('#app')