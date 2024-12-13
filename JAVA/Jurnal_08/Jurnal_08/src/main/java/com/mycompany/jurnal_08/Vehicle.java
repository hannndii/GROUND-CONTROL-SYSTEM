/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.jurnal_08;

/**
 *
 * @author tupog
 */
public abstract class Vehicle {
    private String plateNumber;
    private double capacity;
    
    public Vehicle(String plateNumber, double capacity){
        this.plateNumber = plateNumber;
        this.capacity = capacity;
    }
    
    public abstract double calculateOperationalCost();
    
    void displayInfo(){
        System.out.println("Vehicle Plate: " + this.plateNumber + ", Capacity:" + this.capacity + "kg, Operational Cost:" + this.calculateOperationalCost());          
    }
}
