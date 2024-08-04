import { Image, ScrollView, StyleSheet, Text, View } from "react-native";
import React, { useState } from "react";
import { SafeAreaView } from "react-native-safe-area-context";
import FormField from "../../components/FormField";
import { images } from "../../constants";
import CustotmButtor from "../../components/CustotmButtor";
import { Link } from "expo-router";

const SignUp = () => {
  const [form, setForm] = useState({
    phoneNumber: "",
    email: "",
    password: "",
  });

  const [isSubmitting, setIsSubmitting] = useState(false);

  const submit = () => {};

  return (
    <SafeAreaView className="bg-primary h-full" style={styles.container}>
      <ScrollView>
        <View className="w-full justify-center items-center h-full px-4">
          <Text className="text-3xl text-black font-bold mt-10 ">Create Account</Text>
          <Image source={images.logo} className="mt-5" />
          <FormField
            title="Phone Number"
            value={form.phoneNumber}
            handleChangeText={(e) => setForm({ ...form, phoneNumber: e })}
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
