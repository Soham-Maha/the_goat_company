import { Image, StyleSheet, Text, View } from "react-native";
import { Tabs, Redirect } from "expo-router";

import { icons } from "../../constants";

const TabIcon = ({ icon, color, name, focused }) => {
  return (
    <View style={styles.Icons}>
      <Image
        source={icon}
        resizeMode="contain"
        tintColor={color}
        style={{ height: 27, width: 27 }}
      />
      <Text style={{color: color}}>
        {name}
      </Text>
    </View>
  );
};

const TabsLayout = () => {
  return (
    <>
      <Tabs
        screenOptions={{
          tabBarActiveTintColor: '#000000',
          tabBarInactiveBackgroundColor: '#A4A364',
          tabBarShowLabel: false,
          tabBarStyle: {
            backgroundColor: '#92A286',
            borderTopWidth:1,
            borderTopColor: '#232533',
            height: 60,
          }
        }}
        >
        <Tabs.Screen
          name="home"
          options={{
            title: "Home",
            headerShown: false,
            tabBarIcon: (color, focused) => (
              <TabIcon
                icon={icons.home}
                color={color}
                name="Home"
                focused={focused}
              />
            ),
          }}
        />
        <Tabs.Screen
          name="search"
          options={{
            title: "search",
            headerShown: false,
            tabBarIcon: (color, focused) => (
              <TabIcon
                icon={icons.home}
                color={color}
                name="search"
                focused={focused}
              />
            ),
          }}
        />
        <Tabs.Screen
          name="profile"
          options={{
            title: "Profile",
            headerShown: false,
            tabBarIcon: (color, focused) => (
              <TabIcon
                icon={icons.profile}
                color={color}
                name="Profile"
                focused={focused}
              />
            ),
          }}
        />
      </Tabs>
    </>
  );
};

export default TabsLayout;

const styles = StyleSheet.create({
  Icons:{
    justifyContent: 'center',
    alignItems: 'center'
  }
});
