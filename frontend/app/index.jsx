import { Link, router } from "expo-router";
import { StatusBar } from "expo-status-bar";
import {
  Image,
  ScrollView,
  StyleSheet,
  Text,
  View,
  TouchableOpacity,
} from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { styled } from "nativewind";

import { icons, images } from "../constants";

const StyledView = styled(View);
const StyledText = styled(Text);
const StyledImage = styled(Image);

export default function App() {
  const handlePress = () => {
    router.push("sign-in");
  };
  return (
    <SafeAreaView className="h-full" style={styles.container}>
      <ScrollView contentContainerStyle={{ height: "100%" }}>
        <View className="w-full justify-center items-center h-[95vh] px-4">
        <Image
            source={images.logo}
            className="w-[100px] h-[100px] m-10"
          />
          <StyledView className="flex-1">
            {/* Container for text and boxes */}
            <StyledView className="flex-1 justify-between">
              {/* Text above the boxes */}
              <StyledView className="items-center">
                <StyledText className="text-4xl font-bold text-gray-800">
                  Chose Role
                </StyledText>
              </StyledView>

              {/* Boxes at the bottom */}
              <StyledView className="flex-1 justify-center items-center w-[390px] h-[844px]">
                <StyledView className="flex-row">
                  <TouchableOpacity onPress={handlePress}>
                    <StyledView className="w-[166px] h-[230px] justify-center items-center m-2 rounded-[18px]" style={styles.container1}>
                    <StyledImage 
                      source={images.money}
                      className="w-[90px] h-[90px] justify-center"
                    />
                      <StyledText className="text-black font-bold text-2xl">
                        Investor
                      </StyledText>
                    </StyledView>
                  </TouchableOpacity>
                  <TouchableOpacity onPress={handlePress}>
                    <StyledView className="w-[166px] h-[230px] bg-green-500 justify-center items-center m-2 rounded-[18px]" style={styles.container2}>
                    <StyledImage 
                      source={images.goat}
                      className="w-[90px] h-[90px] justify-center"
                    />
                      <StyledText className="text-black font-bold text-2xl">
                        Farmer
                      </StyledText>
                    </StyledView>
                  </TouchableOpacity>
                </StyledView>
                <StyledView className="flex-row">
                  <TouchableOpacity onPress={handlePress}>
                    <StyledView className="w-[166px] h-[230px] bg-green-500 justify-center items-center m-2 rounded-[18px]" style={styles.container3}>
                    <StyledImage 
                      source={images.buyer}
                      className="w-[90px] h-[90px] justify-center"
                    />
                      <StyledText className="text-black font-bold text-2xl">
                        Buyer
                      </StyledText>
                    </StyledView>
                  </TouchableOpacity>
                  <TouchableOpacity onPress={handlePress}>
                    <StyledView className="w-[166px] h-[230px] bg-green-500 justify-center items-center m-2 rounded-[18px]" style={styles.container4}>
                    <StyledImage 
                      source={images.vet_icon}
                      className="w-[90px] h-[90px] justify-center"
                    />
                      <StyledText className="text-black font-bold text-2xl">
                        Veterinarian
                      </StyledText>
                    </StyledView>
                  </TouchableOpacity>
                </StyledView>
              </StyledView>
            </StyledView>
          </StyledView>
        </View>
      </ScrollView>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container1: {
    backgroundColor: "#89998E",
  },
  container2: {
    backgroundColor: "#93A186",
  },
  container3: {
    backgroundColor: "#929876",
  },
  container4: {
    backgroundColor: "#777B6D",
  },
  container:{
    backgroundColor: "#B99C7C"
  }
});
