import { StyleSheet, Text, TouchableOpacity, View } from 'react-native'
import React from 'react'

const CustotmButton = ({title, handlePress, containerStyles, textStyles, isLoading}) => {
  return (
    <TouchableOpacity 
    onPress={handlePress}
    activeOpacity={0.7}
    disabled={isLoading}
    style={styles.container1} className={`rounded-[67px] min-h-[41px] justify-center items-center ${containerStyles} ${isLoading ? 'opacity:50': ''}`}>
      <Text className={`text-lg ${textStyles}`}>{title}</Text>
    </TouchableOpacity>
  )
}

export default CustotmButton

const styles = StyleSheet.create({
    container1:{
        backgroundColor: '#93A186'
    }
})