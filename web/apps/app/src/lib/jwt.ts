import {set} from "@vueuse/core";

export function useJWT() {
    const token = useLocalStorage('user-token', '')

    const refreshToken = useLocalStorage('refresh-token', '')

    const setToken = (props: string) => {
        set(token, props)
    }

    const setRefreshToken = (props: string) => {
        set(refreshToken, props)
    }
    return {
        setToken,
        token,
        refreshToken,
        setRefreshToken,
    }
}