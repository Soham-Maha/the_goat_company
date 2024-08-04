import { Image, ScrollView, StyleSheet, Text, View } from "react-native";
import React from "react";
import { SafeAreaView } from "react-native-safe-area-context";
import { styled } from "nativewind";
import icons from "../../constants/icons";

const StyledView = styled(View);
const StyledText = styled(Text);
const StyledImage = styled(Image);

const Home = () => {
  return (
    <SafeAreaView className="h-full" style={styles.container}>
      <ScrollView contentContainerStyle={{ height: "100%" }}>
        <StyledView className="justify-between flex flex-row w-full">
          <Text className="text-3xl pl-10 pt-5 pb-5 m-1 font-bold">
            G.O.A.T.
          </Text>
          <Image source={icons.chat} className="w-[30px] h-[30px] p-5 m-5" />
        </StyledView>
        <StyledView className="flex-1 items-center">
          <StyledView className="flex flex-col space-y-4 p-4">
            <StyledView
              className="min-w-full h-[133px] rounded-[15px] justify-center"
              style={styles.container1}
            >
              <StyledText className="justify-center items-center pl-4 text-white font-bold text-[50px] ">
                +$1000
              </StyledText>
            </StyledView>
            <StyledView
              className="min-w-full h-[48px] rounded-[5px] justify-center pl-2"
              style={styles.container2}
            >
              <StyledText className="text-[20px] text-red-700">
                ALERT NOTIFICATION!
              </StyledText>
            </StyledView>
            <StyledView
              className="min-w-full h-[48px] rounded-[5px] justify-center pl-2"
              style={styles.container2}
            >
              <StyledText className="text-[20px] text-red-700">
                ALERT NOTIFICATION!
              </StyledText>
            </StyledView>
          </StyledView>
          <StyledView
            className="w-[363px] h-[50px] rounded-[5px] "
            style={styles.container}
          >
            <StyledText className="font-bold text-2xl text-[30px] text-left ">
              MY ORDERS
            </StyledText>
          </StyledView>
          <StyledView
            className="w-full h-[50px] border-t flex-row justify-between"
            style={styles.container}
          >
            <StyledText className="font-bold text-2xl p-1 text-[25px] text-left">
              Goat-Code
            </StyledText>
            <StyledText className="font-bold text-2xl p-1 text-[25px] text-right">
              Price
            </StyledText>
          </StyledView>
          <StyledView
            className="w-full h-[50px] border-t flex-row justify-between"
            style={styles.container}
          >
            <StyledText className="font-bold text-2xl p-1 text-[25px] text-left">
              Goat-Code
            </StyledText>
            <StyledText className="font-bold text-2xl p-1 text-[25px] text-right">
              Price
            </StyledText>
          </StyledView>
          <StyledView
            className="w-full h-[50px] border-t flex-row justify-between"
            style={styles.container}
          >
            <StyledText className="font-bold text-2xl p-1 text-[25px] text-left">
              Goat-Code
            </StyledText>
            <StyledText className="font-bold text-2xl p-1 text-[25px] text-right">
              Price
            </StyledText>
          </StyledView>
          <StyledView
            className="w-full h-[50px] border-t border-b flex-row justify-between"
            style={styles.container}
          >
            <StyledText className="font-bold text-2xl p-1 text-[25px] text-left">
              Goat-Code
            </StyledText>
            <StyledText className="font-bold text-2xl p-1 text-[25px] text-right">
              Price
            </StyledText>
          </StyledView>
        </StyledView>
      </ScrollView>
    </SafeAreaView>
  );
};

export default Home;

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
