/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.trialobject;

/**
 *
 * @author tupog
 */
public class Taxi extends Car{
    private int maxNum;

    public int getMaxNum() {
        return maxNum;
    }

    public void setMaxNum(int maxNum) {
        this.maxNum = maxNum;
    }
    
    @Override
    public void Go(int distance){
       System.out.println(distance + "km berpindah");
   }
}
