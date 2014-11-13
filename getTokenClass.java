import  java.io.*;


public class getTokenClass {

	public static String line = null;//holds current line for changes in both methods
	
	/*	
	* 	Reads in file called input.txt.
	*	Main then takes in line in a buffer reader and feeds in into get token until line 
	*	until line is only equal to ""
	*/
	public static void main(String[] args) throws Exception{

		String token = "";
		int size = 0;
		try{
			BufferedReader text  = new BufferedReader(new FileReader("input.txt"));
			while ((line = text.readLine()) != null){
				
				while(line.length() > 0){
					token = getToken(line);
					System.out.println(token);

				}

			}
		}catch(Exception e){
			e.printStackTrace();
		}// end of catch try
	}//end of main	
	
	/*	
	*	This method accepts any string and breaks it up returning on token at a time by first
	*	getting a stand alone line or string or number until format programmed is accepted.
	*/
	public static String getToken(String thisLine){

		String token = "";
		token += thisLine.charAt(0);
		int i = 0, size;


		if(thisLine.charAt(i) >= 'A' && thisLine.charAt(i) <= 'Z'){   //token or keyword in caps
			token = "";
			while( i<thisLine.length() && thisLine.charAt(i) >= 'A' && thisLine.charAt(i) <= 'Z' ){
				token += thisLine.charAt(i);
				i++;
			}


		}else if(thisLine.charAt(i) >= '0' && thisLine.charAt(i) <= '9'){//token or keyword in caps
			i++;
			while( i<thisLine.length() && thisLine.charAt(i) >= '0' 
									   && thisLine.charAt(i) <= '9' 
									   || thisLine.charAt(i) == '.' ){
				token += thisLine.charAt(i);
				i++;
			}

            if( token.length() < 7){		//size setter
                int j = 7 - token.length();

				size = token.length();
				line = line.substring(size);	

                while( j > 0){
                	token += " ";
                	j--;
            	}

           		return (token+" 29");
           	}else{
				size = token.length();
				line = line.substring(size);					            	
           		return (token+" 29");
            }



		}else if( thisLine.charAt(0) == ':' && thisLine.charAt(1) == '=') {  // := input
			token += thisLine.charAt(1);


		}else if( thisLine.charAt(0) == '<' && thisLine.charAt(1) == '=') {  // <= input
			token += thisLine.charAt(1);

		}else if( thisLine.charAt(0) == '<' && thisLine.charAt(1) == '>') {  // <> input
			token += thisLine.charAt(1);

		}else if( thisLine.charAt(0) == '>' && thisLine.charAt(1) == '=') {  // >= input
			token += thisLine.charAt(1);


		}else if(token.equals("\"")){    //string inside " "  
            i ++;
        	token = "";
        	String mark = "";
           	for(i=i;  !mark.equals("\""); i++){
          		mark = "";
           		token += thisLine.charAt(i);
           		mark += thisLine.charAt(i+1);
          	}
            if( token.length() < 7){
                int j = 7 - token.length();

				size = token.length();
				line = line.substring(size+2);	

                while( j > 0){
                	token += " ";
                	j--;
            	}

           		return (token+" 30");
           	}else{
				size = token.length();
				line = line.substring(size+2);					            	
           		return (token+" 30");
            }
        }

		size = token.length();			
		line = line.substring(size);	//changing line
		
		if( line.length() == 0  ){		//End of line check
			line = "EOLN";
		}
		
		switch (token) {//generic
   			
   			case "IF":		
	  			return "IF       1";
	        case "THEN":
       			return "THEN     2";	    	   	
       		case "ELSE":
      			return "ELSE     3";           		
        	case "FI":
        		return "FI       4";               	
        	case "LOOP":
        		return "LOOP     5";        	
        	case "BREAK":
        		return "BREAK    6";               	
           	case "READ":
           		return "READ     7";               	
        	case "PRINT":
           		return "PRINT    8";               	
        	case "AND":
          		return "AND      9";               	
        	case "OR":
           		return "OR      10";                
        	case ".":
           		return ".       11";               	
        	case ")":
           		return ")       12";               	
        	case "(":
          		return "(       13";               	
        	case "/":
           		return "/       14";               	
         	case "*":
         		return "*       15";               	
         	case "-":
           		return "-       16";               	
        	case "+":
           		return "+       17";              	
        	case "<>":
           		return "<>      18";               	
        	case ">":
           		return ">       19";               	
        	case ">=":
           		return ">=      20";               	
    	   	case "=":
           		return "=       21";              	
        	case "<=":
          		return "<=      22";              	
        	case "<":
           		return "<       23";               	
        	case ":=":
           		return ":=      24";
           	case ";":
           		return ";       25";
            case " ":
                return "SPACE   26";
            case "EOLN":
            	line = "";
                return "EOLN    27";
        	default: 
                if( token.length() < 7){
            	    int j = 7 - token.length();
                    while( j > 0){
                   	    token += " ";
                        j--;
                   	}
        	       	token += " 28";
            	    return token;
                }else{
                   	token += " 28";
          		    return token;
                }
            }//end of switch;
	}//end of getToken
}//end of class	

