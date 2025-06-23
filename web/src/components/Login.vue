<!-- <template>
    <div class="form-container">
        <div class="form">
            <form @submit.prevent="submit">
                <h1 class="h3 mb-3 fw-normal" />

                <div v-if="!tokenRequired" class="form-floating">
                    <input id="floatingInput" v-model="username" type="text" class="form-control" placeholder="Username" autocomplete="username" required>
                    <label for="floatingInput">{{ $t("Username") }}</label>
                </div>

                <div v-if="!tokenRequired" class="form-floating mt-3">
                    <input id="floatingPassword" v-model="password" type="password" class="form-control" placeholder="Password" autocomplete="current-password" required>
                    <label for="floatingPassword">{{ $t("Password") }}</label>
                </div>

                <div v-if="tokenRequired">
                    <div class="form-floating mt-3">
                        <input id="otp" v-model="token" type="text" maxlength="6" class="form-control" placeholder="123456" autocomplete="one-time-code" required>
                        <label for="otp">{{ $t("Token") }}</label>
                    </div>
                </div>

                <div class="form-check mb-3 mt-3 d-flex justify-content-center pe-4">
                    <div class="form-check">
                        <input id="remember" v-model="$root.remember" type="checkbox" value="remember-me" class="form-check-input">

                        <label class="form-check-label" for="remember">
                            {{ $t("Remember me") }}
                        </label>
                    </div>
                </div>
                <button class="w-100 btn btn-primary" type="submit" :disabled="processing">
                    {{ $t("Login") }}
                    "Login"
                </button>

                <div v-if="res && !res.ok" class="alert alert-danger mt-3" role="alert">
                    {{ $t(res.msg) }}
                    {{ res.msg }}
                </div>
            </form>
        </div>
    </div>
</template> -->
<template>
    <div class="w-full h-full">
        <div class="h-50% mx-auto w-18% flex flex-col justify-center items-center bg-violet rounded-2xl">
            <img src="@/assets/meow.png" class="w-50px h-50px rounded-50% mb-1">
            <p class="mb-5">meow</p>
            <el-form :model="form">
                <el-form-item>
                    <el-input v-model="form.username" size="large" clearable>
                        <template #prefix>
                            <i-ep-user class="text-violet"></i-ep-user>
                        </template>
                    </el-input>
                </el-form-item>
                <el-form-item>
                    <el-input v-model="form.password" type="password" show-password size="large" clearable>
                        <template #prefix>
                            <i-ep-lock class="text-violet"></i-ep-lock>
                        </template>
                    </el-input>
                </el-form-item>
                <el-row justify="end">
                    <el-form-item>
                        <el-button type="primary" @click="submit">
                            登录
                        </el-button>
                    </el-form-item>
                </el-row>
            </el-form>
        </div>
    </div>
</template>

<script setup lang="ts">

interface Form {
    username: string
    password: string
}


const form = ref<Form>({} as Form)

const processing = ref(false)
const res = ref(null)
// const token = ref("")
// const tokenRequired = ref(false)


onMounted(() => {
    document.title += " - Login"
})

onBeforeUnmount(() => {
    document.title = document.title.replace(" - Login", "");
})

const submit = () => {
    processing.value = true
    // login && need token
}

</script>

<style lang="scss" scoped>
.form-container {
    display: flex;
    align-items: center;
    padding-top: 40px;
    padding-bottom: 40px;
}

.form-floating {
    >label {
        padding-left: 1.3rem;
    }

    >.form-control {
        padding-left: 1.3rem;
    }
}

.form {
    width: 100%;
    max-width: 330px;
    padding: 15px;
    margin: auto;
    text-align: center;
}
</style>
