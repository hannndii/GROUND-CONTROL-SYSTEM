/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.jurnal_08;

/**
 *
 * @author tupog
 */
public class Truck extends Vehicle{
    private double fuelConsumtion;
    private double distance;
    
    public Truck(String plateNumber, double capacity, double fuelConsumtion, double distance){
        super(plateNumber, capacity);
        this.fuelConsumtion = fuelConsumtion;
        this.distance = distance;
    }
    
    @Override
    public double calculateOperationalCost(){
        return this.fuelConsumtion * this.distance * 1.5;
    }
    
}
