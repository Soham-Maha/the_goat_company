import { Image, StyleSheet, Text, View } from "react-native";
import React from "react";
import { SafeAreaView } from "react-native-safe-area-context";
import { styled } from "nativewind";
import { icons } from "../../constants";

const StyledView = styled(View);
const StyledText = styled(Text);
const StyledImage = styled(Image);

const Profile = () => {
  return (
    <SafeAreaView className="h-full" style={styles.container}>
      <StyledView className="justify-between flex flex-row w-full">
        <Text className="text-3xl pl-10 pt-5 pb-5 m-1 font-bold">G.O.A.T.</Text>
        <Image source={icons.chat} className="w-[30px] h-[30px] p-5 m-5" />
      </StyledView>
      <View className="flex-1 relative items-center">
        <View
          className="p-4 rounded-[15px] w-[259px] h-[274px] items-center"
          style={styles.container1}
        >
          <Image
            source={icons.edit}
            className="absolute right-[10px] top-2 w-[22px] h-[22px]"
          />
          <Text className="text-white"></Text>
          <View className="flex-1  items-center">
            <View
              className="rounded-lg w-[152px] h-[149px] justify-center items-center"
              style={styles.container2}
            >
              <Image
                source={icons.Group}
                className="justify-center items-center w-[78px] h-[78px]"
              />
            </View>
            <Text className="text-2xl">Satish Jain</Text>
            <Text className="text-xl">(farmercode)</Text>
          </View>
        </View>
        <StyledView
            className="w-[363px] h-[50px] rounded-[5px] "
            style={styles.container}
          >
            <StyledText className="font-bold text-2xl text-[30px] text-left ">
              Funds
            </StyledText>
          </StyledView>
          <StyledView
            className="w-full h-[50px] border-t flex-row justify-between border-b" 
            style={styles.container}
          >
            <StyledText className="font-bold text-2xl p-1 text-[25px] text-left">
              Goat-Code
            </StyledText>
            <StyledText className="font-bold text-2xl p-1 text-[25px] text-right">
              Price
            </StyledText>
          </StyledView>
          
      </View>
      <StyledView>
      </StyledView>
    </SafeAreaView>
  );
};

export default Profile;

const styles = StyleSheet.create({
  container: {
    backgroundColor: "#B99C7C",
  },
  container1: {
    backgroundColor: "#D0D8C1",
  },
  container2: {
    backgroundColor: "#686464",
  },
});
