import { Image, ScrollView, StyleSheet, Text, View } from "react-native";
import React, { useState } from "react";
import { SafeAreaView } from "react-native-safe-area-context";
import FormField from "../../components/FormField";
import { images } from "../../constants";
import CustotmButtor from "../../components/CustotmButtor";
import { Link } from "expo-router";

const SignIn = () => {
  const [form, setForm] = useState({
    email: "",
    password: "",
  });

  const [isSubmitting, setIsSubmitting] = useState(false);

  const submit = () => {};

  return (
    <SafeAreaView className="bg-primary h-full">
      <ScrollView>
        <View className="w-full justify-center items-center h-full px-4">
          <Text className="text-3xl text-black font-bold mt-10 ">Sign In</Text>
          <Image source={images.logo} className="mt-5" />
          <FormField
            title="Email"
            value={form.email}
            handleChangeText={(e) => setForm({ ...form, email: e })}
            otherStyles="mt-4"
            keyboardType="email-adress"
          />
          <FormField
            title="Password"
            value={form.password}
            handleChangeText={(p) => setForm({ ...form, password: p })}
            otherStyles="mt-4"
          />
          <CustotmButtor
            title="Sign In"
            handlePress={submit}
            containerStyles="w-[115px] h-[41px] mt-5 "
            isLoading={isSubmitting}
          />
          <View className="justify-center flex-row gap-2 pt-5 ">
            <Text className="text-lg text-black">Dont have a account?</Text>
            <Link className="font-bold text-lg text-black" href='sign-up'>Sign up</Link>
          </View>
        </View>
      </ScrollView>
    </SafeAreaView>
  );
};

export default SignIn;

const styles = StyleSheet.create({});
