/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.jurnal_08;

/**
 *
 * @author tupog
 */
public class HeavyDutyTruck extends Vehicle{
    private double fuelConsumtion;
    private double distance;
    
    public HeavyDutyTruck(String plateNumber, double capacity, double distance, double fuelConsumtion){
        super(plateNumber, capacity);
        this.distance = distance;
        this.fuelConsumtion = fuelConsumtion;
    }
    
    @Override
    public double calculateOperationalCost(){
        return this.fuelConsumtion * this.distance * 1.5;
    }
}