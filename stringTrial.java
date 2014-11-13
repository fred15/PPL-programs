
public class stringTrial {

	public static String c = "";

	public static void main(String[] args) {

		c = "talk";

		String talk = "talk";

		talk = changeString(talk);
		System.out.println(talk);
		System.out.println(c);

	}

	public static String changeString( String talk){
		c = "hi";
		System.out.println(c.length());
		return "";
	}
}
