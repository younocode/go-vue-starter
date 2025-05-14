<script setup lang="ts">
import {Button} from '@/components/ui/button'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import {Input} from '@/components/ui/input'
import {Label} from '@/components/ui/label'
import {
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
} from '@/components/ui/tabs'
import {set} from "@vueuse/core";
import {useApiFetch} from "@/lib/apiFetch.ts";
import router from "@/router";

const rememberMe = useLocalStorage('user-remember-store', false)
const signInStateStore = useLocalStorage('user-sign-store', {
  email: '',
  password: '',
})
const signInState = ref(toValue(rememberMe) ? {...toValue(signInStateStore)} : {
  email: '',
  password: '',
})
const signUpState = ref({
  email: '',
  password: '',
  confirmPassword: ''
})

watchDeep(signInState, () => {
  if (toValue(rememberMe)) {
    set(signInStateStore, toValue(signInState))
  }
})

watch(rememberMe, () => {
  if (!toValue(rememberMe)) {
    set(signInStateStore, {
      email: '',
      password: '',
    })
  } else {
    set(signInStateStore, toValue(signInState))
  }
})

function login() {
  useApiFetch('/login', {
    method: 'POST',
    body: JSON.stringify(toValue(signInState))
  })
}

function register() {
  useApiFetch('/sign-up', {
    method: 'POST',
    body: JSON.stringify(toValue(signUpState))
  })
}

function resetPasswd() {
  router.push({
    name: 'resetPassword',
    params: {
      email: 'test',
    }
  })
}
</script>

<template>
 <div class=" h-svh flex items-center justify-center">
   <Tabs default-value="login" class="w-[400px]">
     <TabsList class="grid w-full grid-cols-2">
       <TabsTrigger value="login">
         Login
       </TabsTrigger>
       <TabsTrigger value="register">
         Register
       </TabsTrigger>
     </TabsList>
     <TabsContent value="login">
       <Card>
         <CardHeader>
           <CardTitle>Login</CardTitle>
           <CardDescription>
             Enter your email below to login to your account
           </CardDescription>
         </CardHeader>
         <CardContent class="space-y-2">
           <div class="space-y-1">
             <Label for="email">Email</Label>
             <Input id="email" v-model="signInState.email" placeholder="name@example.com" type="email"/>
           </div>
           <div class="space-y-1">
             <div class="flex items-center">
               <Label for="password">Password</Label>
               <a class="ml-auto inline-block text-sm underline" href="#" @click="resetPasswd">Forgot your password?</a>
             </div>

             <Input id="password" v-model="signInState.password" placeholder="******" type="password"/>
           </div>
           <div class="flex items-center space-x-2">
             <Checkbox id="remember" v-model="rememberMe"/>
             <label
                 for="remember"
                 class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
             >
               Remember Me
             </label>
           </div>
         </CardContent>
         <CardFooter>
           <Button class="w-full" @click="login">Login</Button>
         </CardFooter>
       </Card>
     </TabsContent>
     <TabsContent value="register">
       <Card>
         <CardHeader>
           <CardTitle>Sign Up</CardTitle>
           <CardDescription>
             Enter your information to create an account
           </CardDescription>
         </CardHeader>
         <CardContent class="space-y-2">
           <div class="space-y-1">
             <Label for="email">Email</Label>
             <Input id="email" v-model="signUpState.email" placeholder="name@example.com" type="email"/>
           </div>
           <div class="space-y-1">
             <Label for="password">Password</Label>
             <Input id="password" v-model="signUpState.password" placeholder="******" type="password"/>
           </div>
           <div class="space-y-1">
             <Label for="confirmPassword">Password</Label>
             <Input id="confirmPassword" v-model="signUpState.confirmPassword" placeholder="******" type="password"/>
           </div>
         </CardContent>
         <CardFooter>
           <Button class="w-full" @click="register">Create an account</Button>
         </CardFooter>
       </Card>
     </TabsContent>
   </Tabs>
 </div>
</template>

<style scoped>

</style>
