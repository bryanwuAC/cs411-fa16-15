package csv;

import java.io.*;
import java.util.*;
import javax.xml.parsers.*;
import org.xml.sax.*;
import org.xml.sax.helpers.*;

public class Parser {

	private final int maxAuthorsPerPaper = 200;

	private class ConfigHandler extends DefaultHandler {

		private Locator locator;

		private String personName;
		private String title;
		private String url;
		private int year;
		private String key;
		private String recordTag;
		private Person[] persons = new Person[maxAuthorsPerPaper];
		private int numberOfPersons = 0;

		private boolean insidePerson;
		private boolean insideTitle;
		private boolean insideYear;
		private boolean insideUrl;
		
		public void setDocumentLocator(Locator locator) {
			this.locator = locator;
		}

		public void startElement(String namespaceURI, String localName, String rawName, Attributes atts)
				throws SAXException {
			String k;

			if (insidePerson = (rawName.equals("author") || rawName.equals("editor"))) {
				personName = "";
				return;
			}

			insideTitle = rawName.equals("title");
			insideYear = rawName.equals("year");
			insideUrl = rawName.equals("ee"); // The url in xml is local to dblp
			
			if ((atts.getLength() > 0) && ((k = atts.getValue("key")) != null)) {
				key = k;
				recordTag = rawName;
				title = "";
				year = Integer.MIN_VALUE;
				url = "";
			}
		}

		public void endElement(String namespaceURI, String localName, String rawName) throws SAXException {

			if (rawName.equals("author") || rawName.equals("editor")) {

				Person p;
				if ((p = Person.searchPerson(personName)) == null) {
					p = new Person(personName);
				}
				p.increment();
				if (numberOfPersons < maxAuthorsPerPaper)
					persons[numberOfPersons++] = p;
				return;
			}

			if (rawName.equals(recordTag)) {
				if (numberOfPersons == 0)
					return;
				Person pa[] = new Person[numberOfPersons];
				for (int i = 0; i < numberOfPersons; i++) {
					pa[i] = persons[i];
					persons[i] = null;
				}
				Publication p = new Publication(key, title, pa, year, url);
				numberOfPersons = 0;
			}

		}

		public void characters(char[] ch, int start, int length) throws SAXException {
			if (insidePerson) {
				personName += new String(ch, start, length);
			}

			if (insideTitle) {
				title += new String(ch, start, length);
			}

			if (insideYear) {
				year = Integer.parseInt(new String(ch, start, length));
			}
			
			if (insideUrl) {
				url += new String(ch, start, length);
			}
		}

		private void Message(String mode, SAXParseException exception) {
			System.out.println(mode + " Line: " + exception.getLineNumber() + " URI: " + exception.getSystemId() + "\n"
					+ " Message: " + exception.getMessage());
		}

		public void warning(SAXParseException exception) throws SAXException {

			Message("**Parsing Warning**\n", exception);
			throw new SAXException("Warning encountered");
		}

		public void error(SAXParseException exception) throws SAXException {

			Message("**Parsing Error**\n", exception);
			throw new SAXException("Error encountered");
		}

		public void fatalError(SAXParseException exception) throws SAXException {

			Message("**Parsing Fatal Error**\n", exception);
			throw new SAXException("Fatal Error encountered");
		}
	}

	private void nameLengthStatistics() {
		Iterator i = Person.iterator();
		Person p;
		int l = Person.getMaxNameLength();
		int lengthTable[] = new int[l + 1];
		int j;

		System.out.println();
		System.out.println("Name length: Number of persons");
		while (i.hasNext()) {
			p = (Person) i.next();
			lengthTable[p.getName().length()]++;
		}
		for (j = 1; j <= l; j++) {
			System.out.print(j + ": " + lengthTable[j] + "  ");
			if (j % 5 == 0)
				System.out.println();
		}
		System.out.println();
	}

	private void publicationCountStatistics() {
		Iterator i = Person.iterator();
		Person p;
		int l = Person.getMaxPublCount();
		int countTable[] = new int[l + 1];
		int j, n;

		System.out.println();
		System.out.println("Number of publications: Number of persons");
		while (i.hasNext()) {
			p = (Person) i.next();
			countTable[p.getCount()]++;
		}
		n = 0;
		for (j = 1; j <= l; j++) {
			if (countTable[j] == 0)
				continue;
			n++;
			System.out.print(j + ": " + countTable[j] + "  ");
			if (n % 5 == 0)
				System.out.println();
		}
		System.out.println();
	}

	Parser(String uri) {
		try {
			SAXParserFactory parserFactory = SAXParserFactory.newInstance();
			SAXParser parser = parserFactory.newSAXParser();
			ConfigHandler handler = new ConfigHandler();
			parser.getXMLReader().setFeature("http://xml.org/sax/features/validation", true);
			parser.parse(new File(uri), handler);
		} catch (IOException e) {
			System.out.println("Error reading URI: " + e.getMessage());
		} catch (SAXException e) {
			System.out.println("Error in parsing: " + e.getMessage());
		} catch (ParserConfigurationException e) {
			System.out.println("Error in XML parser configuration: " + e.getMessage());
		}
		System.out.println("Number of Persons : " + Person.numberOfPersons());
//		nameLengthStatistics();
		System.out.println("Number of Publications with authors/editors: " + Publication.getNumberOfPublications());
		System.out
				.println("Maximum number of authors/editors in a publication: " + Publication.getMaxNumberOfAuthors());
//		publicationCountStatistics();
		Person.enterPublications();
	}

	public static void printCSV() throws FileNotFoundException, UnsupportedEncodingException {
		Iterator paperIterator = Publication.iterator();
		Iterator personIterator = Person.iterator();

		PrintWriter paperWriter = new PrintWriter("paper.csv", "ISO-8859-1");
		PrintWriter writtenWriter = new PrintWriter("written.csv", "ISO-8859-1");
		while (paperIterator.hasNext()) {
			Publication paper = (Publication) paperIterator.next();
			
			if (paper.getYear() != Integer.MIN_VALUE) {
				paperWriter.println("\"" + paper.getKey() + "\",\"" + paper.getTitle() + "\"," + paper.getYear() + ",\"" + paper.getUrl() + "\"");
			} else {
				paperWriter.println("\"" + paper.getKey() + "\",\"" + paper.getTitle() + "\",,\"" + paper.getUrl() + "\"");
			}
			
			for (Person person : paper.getAuthors()) {
				writtenWriter.println(paper.getKey() + "," + person.getKey());
			}
		}
		paperWriter.close();
		writtenWriter.close();

		PrintWriter personWriter = new PrintWriter("person.csv", "ISO-8859-1");
		while (personIterator.hasNext()) {
			Person person = (Person) personIterator.next();
			personWriter.println(person.getKey() + ",\"" + person.getName() + "\"");
		}
		personWriter.close();

	}

	public static void main(String[] args) throws FileNotFoundException, UnsupportedEncodingException {
		if (args.length < 1) {
			System.out.println("Usage: java Parser [input]");
			System.exit(0);
		}
		new Parser(args[0]);
		System.out.println("Max Title Length : " + Publication.getMaxTitleLength());
		System.out.println("Max URL Length : " + Publication.getMaxUrlLength());
		System.out.println("Max Key Length : " + Publication.getMaxKeyLength());

	}

}
