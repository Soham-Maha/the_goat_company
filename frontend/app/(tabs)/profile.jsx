import { Image, StyleSheet, Text, View } from 'react-native'
import React from 'react'
import { SafeAreaView } from 'react-native-safe-area-context'
import { styled } from "nativewind";

const StyledView = styled(View);
const StyledText = styled(Text);
const StyledImage = styled(Image);

const Profile = () => {
  return (
    <SafeAreaView className="h-full" style={styles.container}>
      <StyledView className="justify-between flex flex-row w-full">
          <Text className="text-3xl pl-10 pt-5 pb-5 m-1 font-bold">
            G.O.A.T.
          </Text>
          <Image source={icons.chat} className="w-[30px] h-[30px] p-5 m-5" />
        </StyledView><Text>Profile</Text>
    </SafeAreaView>
  )
}

export default Profile

const styles = StyleSheet.create({
  container: {
    backgroundColor: "#B99C7C",
  },
  container1: {
    backgroundColor: "#4A6651",
  },
  container2: {
    backgroundColor: "#A48364",
  },
});