package csv;

/*
 * Created on 07.06.2005
 */

import java.util.*;

/**
 * @author ley
 *
 * created first in project xml5_coauthor_graph
 */
public class Publication {
    private static Set ps= new HashSet(650000);
    private static int maxNumberOfAuthors = 0;
    private static int maxTitleLength = 0;
    private String key;
    private Person[] authors;	// or editors
	private String title;
	private int year;
	private String url;
    
    public Publication(String key, String title, Person[] persons, int year, String url) {
        this.key = key;
        authors = persons;
        this.title = title;
        this.year = year;
        this.url = url;
        
        ps.add(this);
        if (persons.length > maxNumberOfAuthors)
            maxNumberOfAuthors = persons.length;
    }
    
    public static int getNumberOfPublications() {
        return ps.size();
    }
    
    public static int getMaxNumberOfAuthors() {
        return maxNumberOfAuthors;
    }
    
    public Person[] getAuthors() {
        return authors;
    }
    
    public String getKey(){
    	return key;
    }
    
    public String getTitle(){
    	return title;
    }
    
    public int getYear(){
    	return year;
    }
    
    public String getUrl(){
    	return url;
    }
    
    static Iterator iterator() {
        return ps.iterator();
    }
}
