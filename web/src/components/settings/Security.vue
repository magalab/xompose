<template>
    <div class="w-full h-500px">
        <div class="my-4">
            <!-- Change Password -->
            <p>
                {{ $t("Current User") }}: <strong> xiaolua </strong>
            </p>

            <h5 class="my-4 settings-subheading">{{ $t("Change Password") }}</h5>
            <el-form class="mb-3">
                <el-form-item class="mb-3">
                    <label for="current-password" class="form-label">
                        {{ $t("Current Password") }}
                    </label>
                    <el-input v-model="formRef.currentPassword" type="password" show-password />
                </el-form-item>

                <el-form-item class="mb-3">
                    <label for="new-password" class="form-label">
                        {{ $t("New Password") }}
                    </label>
                    <el-input v-model="formRef.password" type="password" show-password />
                </el-form-item>
                <el-form-item class="mb-3">
                    <label for="repeat-new-password" class="form-label">
                        {{ $t("Repeat New Password") }}
                    </label>
                    <el-input v-model="formRef.password2" type="password" show-password />
                    <div class="color-red" v-if="invalidPassword">
                        {{ $t("PasswordNotMatchMsg") }}
                    </div>
                </el-form-item>
                <el-button type="primary" @click="savePassword">
                    {{ $t("Update Password") }}
                </el-button>
            </el-form>
        </div>
    </div>
</template>

<script setup lang="ts">
import { changePasswordAPI } from "@/api/user";
const invalidPassword = ref(false)

const {t} = useI18n()
interface Form {
    currentPassword: string
    password: string
    password2: string
}

const formRef = ref<Form>({} as Form)

const savePassword = async () => {
    console.log(formRef.value, 123)
    if (formRef.value.password !== formRef.value.password2) {
        invalidPassword.value = true;
        return
    } else {
        await changePasswordAPI({
            currentPassword: formRef.value.currentPassword,
            password: formRef.value.password,
            password2: formRef.value.password2,
        })
        ElMessage.success(t("PasswordUpdatedSuccessfully"))
        formRef.value.currentPassword = ""
        formRef.value.password = ""
        formRef.value.password2 = ""
    }
}
</script>
