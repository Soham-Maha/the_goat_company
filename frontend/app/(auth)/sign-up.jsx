import { Image, ScrollView, StyleSheet, Text, View } from "react-native";
import React, { useState } from "react";
import { SafeAreaView } from "react-native-safe-area-context";
import FormField from "../../components/FormField";
import { images } from "../../constants";
import CustotmButton from "../../components/CustotmButton";
import { Link, router } from "expo-router";

const SignUp = () => {
  const [form, setForm] = useState({
    phoneNumber: "",
    email: "",
    password: "",
  });

  const [isSubmitting, setIsSubmitting] = useState(false);

  const submit = async () => {
    setIsSubmitting(true);
    try {
      const response = await fetch("https://36d8-152-67-176-76.ngrok-free.app/user/signup", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          usertype: "farmer",
          name: form.name,
          email: form.email,
          password: form.password,
        }),
      });

      if (!response.ok) {
        // Handle error response
        console.error("SignUp failed");
        return;
      }

      const data = await response.json();
      console.log("SignUp successful:", data);
      router.push("/sign-in");
    } catch (error) {
      console.error("Error:", error);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <SafeAreaView className="bg-primary h-full" style={styles.container}>
      <ScrollView>
        <View className="w-full justify-center items-center h-full px-4">
          <Text className="text-3xl text-black font-bold mt-10 ">Create Account</Text>
          <Image source={images.logo} className="mt-5" />
          <FormField
            title="Name"
            value={form.name}
            handleChangeText={(e) => setForm({ ...form, name: e })}
            otherStyles="mt-4" 
          />
          <FormField
            title="Email"
            value={form.email}
            handleChangeText={(e) => setForm({ ...form, email: e })}
            otherStyles="mt-4"
            keyboardType="email-address"
          />
          <FormField
            title="Password"
            value={form.password}
            handleChangeText={(p) => setForm({ ...form, password: p })}
            otherStyles="mt-4"
          />
          <CustotmButton
            title="Create Account"
            handlePress={submit}
            containerStyles="w-[200px] h-[77px] mt-5 "
            isLoading={isSubmitting}
          />
          <View className="justify-center flex-row gap-2 pt-5 ">
            <Text className="text-lg text-black">Have an account already?</Text>
            <Link className="font-bold text-lg text-black" href='sign-in'>Login</Link>
          </View>
        </View>
      </ScrollView>
    </SafeAreaView>
  );
};

export default SignUp;

const styles = StyleSheet.create({
  container:{
    backgroundColor: "#B99C7C"
  }
});
