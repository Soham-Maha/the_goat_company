import { StyleSheet, Text, View } from 'react-native'
import React from 'react'
import {Link, Stack} from 'expo-router'

const RootLayout = () => {
  return (
    <Stack>
        <Stack.Screen name='index' options={{headerShown:false}} />
        <Stack.Screen name='(auth)' options={{headerShown:false}} />
        <Stack.Screen name='(tabs)' options={{headerShown:false}} />
        {/* <Stack.Screen name='/search/[query]' options={{headerShown:false}} /> */}
    </Stack>
  )
}

export default RootLayout

const styles = StyleSheet.create({
    container: {
      flex: 1,
      backgroundColor: '#fff',
      alignItems: 'center',
      justifyContent: 'center',
    },
  });