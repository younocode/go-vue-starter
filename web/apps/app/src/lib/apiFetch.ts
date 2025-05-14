import {createFetch} from "@vueuse/core";
import {useJWT} from "@/lib/jwt.ts";


export const useApiFetch = createFetch({
    baseUrl: '/api',
    combination: 'overwrite',
    options: {
        // beforeFetch in pre-configured instance will only run when the newly spawned instance do not pass beforeFetch
        async beforeFetch({ options }) {
            const {token} =  useJWT()
            if (!options.headers ) {
                options.headers = {}
            }
            if(!(<Record<string, string>>options.headers)['Content-Type']) {
                options.headers = {
                    'Content-Type': 'application/json'
                }
            }
            if(!(<Record<string, string>>options)['method']) {
                options.method = "POST"
            }
            (<Record<string, string>>options.headers).Authorization = `Bearer ${toValue(token)}`
            return { options }
        },
    },
})