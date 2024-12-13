/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 */

package com.mycompany.collection;

import java.util.*;

public class Collection {
    public static void main(String[] args) {
        Set set = new HashSet();
        set.add("Bernadine");
        set.add("Elizabeth");
        set.add("Gene");
        set.add("Elizabeth"); // Duplikasi
        set.add("Clara");

        System.out.print("Elemen pada HashSet : ");
        System.out.println(set);

        Set sortSet = new TreeSet(set);
        System.out.print("Elemen pada TreeSet : ");
        System.out.println(sortSet);
    }
}
