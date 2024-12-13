/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.trialobject;

/**
 *
 * @author tupog
 */
public class Car implements Washer{
   private int oilSize;

    public int getOilSize() {
        return oilSize;
    }

    public void setOilSize(int oilSize) {
        this.oilSize = oilSize;
    }
   
   public void Go(int distance){
       System.out.println(distance + "km berpindah");
   }
   
   @Override
   public String washer(){
       return "Mencuci Mobil";
   }
}

