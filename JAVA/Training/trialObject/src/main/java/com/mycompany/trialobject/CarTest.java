/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 */

package com.mycompany.trialobject;
/**
 *
 * @author tupog
 */
public class CarTest {

    public static void main(String[] args) {
        Car car = new Car();
        Car truck = new Truck();
        Car taxi = new Taxi();
        
        System.out.println(car.washer());
        System.out.println("Taksi:" + taxi.washer());
        System.out.println("Truck:" + truck.washer());
    }    
    
}
