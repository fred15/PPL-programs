

public class arrayman {
	public static void main(String[] args){
		
		String line = "\"this is me\"";
		char[] cArray = line.toCharArray();
		line = "";
		
		line += cArray[0];

		if( line.equals("\"")){
			System.out.println("yay");
		}
	}
}