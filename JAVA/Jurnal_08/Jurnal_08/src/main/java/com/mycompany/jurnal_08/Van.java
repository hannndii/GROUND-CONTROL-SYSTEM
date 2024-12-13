/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.jurnal_08;

/**
 *
 * @author tupog
 */
public class Van extends Vehicle{
    private double hourlyRate;
    private int hoursUsed;
    
    public Van(String plateNumber, double capacity, double hourlyRate, int hoursUsed){
        super(plateNumber, capacity);
        this.hourlyRate = hourlyRate;
        this.hoursUsed = hoursUsed;
    }
    
    @Override
    public double calculateOperationalCost(){
        return this.hourlyRate * this.hoursUsed;
    }
}
