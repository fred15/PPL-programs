import  java.io.*;


public class getThis {
	public static void main(String[] args) throws Exception{

		String line = null, token = "";
		char[] cArray; 
		try{
			BufferedReader text  = new BufferedReader(new FileReader("token.txt"));
			while ((line = text.readLine()) != null){
				cArray = line.toCharArray();

				
				for(int i = 0; i<cArray.length; i++){
					if(cArray[i] != ' '){
						token += cArray[i];

						if(token.equals("\"")){    //string inside " "
                    		i ++;
                    		token = "";
                    		String mark = "";
            	        	for(i=i;  !mark.equals("\""); i++){
            	        		mark = "";
            	        		token += cArray[i];
            	        		mark += cArray[i+1];
            	        	}
                            if( token.length() < 7){
                                int j = 7 - token.length();
                                while( j > 0){
                                    token += " ";
                                    j--;
                             	}
                                System.out.println(token+" 30");
                            }else{
                                System.out.println(token+" 30");
                            } 
                  			token = "";
                		}
					
					}else if( isDouble(token) ){	//numbers
                        if( token.length() < 7){
                            int j = 7 - token.length();
                            while( j > 0){
                                token += " ";
                                j--;
                            }
                            System.out.println(token+" 29");
                        }else{
                            System.out.println(token+" 29");
                        }  
                     	System.out.println("SPACE   26");
                        token = "";

                    }else if(token != ""){

						switch (token) {//generic
            				case "IF":		
            					System.out.println("IF       1");
                   				System.out.println("SPACE   26");
               		    		break;
            				case "THEN":
            					System.out.println("THEN     2");
                     			System.out.println("SPACE   26");
                 			   	break;
            				case "ELSE":
            					System.out.println("ELSE     3");
                     			System.out.println("SPACE   26");
                     			break;
          					case "FI":
            					System.out.println("FI       4");
                     			System.out.println("SPACE   26");
                     			break;
            				case "LOOP":
            					System.out.println("LOOP     5");
                     			System.out.println("SPACE   26");
                     			break;
            				case "BREAK":
            					System.out.println("BREAK    6");
                     			System.out.println("SPACE   26");
                     			break;
            				case "READ":
            					System.out.println("READ     7");
                     			System.out.println("SPACE   26");
                     			break;
        			    	case "PRINT":
            					System.out.println("PRINT    8");
                     			System.out.println("SPACE   26");
                     			break;
         			  		case "AND":
            					System.out.println("AND      9");
                     			System.out.println("SPACE   26");
                     			break;
          			 		case "OR":
            					System.out.println("OR      10");
                     			System.out.println("SPACE   26");
                     			break;
          				 	case ".":
            					System.out.println(".       11");
                     			System.out.println("SPACE   26");
                     			break;
         				   	case ")":
            					System.out.println(")       12");
                     			System.out.println("SPACE   26");
                     			break;
         			   		case "(":
            					System.out.println("(       13");
                     			System.out.println("SPACE   26");
                     			break;
         			   		case "/":
            					System.out.println("/       14");
                     			System.out.println("SPACE   26");
                     			break;
        	 			   	case "*":
            					System.out.println("*       15");
                     			System.out.println("SPACE   26");
                     			break;
        	 			   	case "-":
            					System.out.println("-       16");
                     			System.out.println("SPACE   26");
                     			break;
         				   	case "+":
            					System.out.println("+       17");
                     			System.out.println("SPACE   26");
                     			break;
         				   	case "<>":
            					System.out.println("<>      18");
                     			System.out.println("SPACE   26");
                     			break;
         				   	case ">":
            					System.out.println(">       19");
                     			System.out.println("SPACE   26");
                     			break;
         			   		case ">=":
            					System.out.println(">=      20");
                     			System.out.println("SPACE   26");
                     			break;
    	     			   	case "=":
            					System.out.println("=       21");
                     			System.out.println("SPACE   26");
                     			break;
        	 			   	case "<=":
            					System.out.println("<=      22");
                     			System.out.println("SPACE   26");
                     			break;
         				   	case "<":
            					System.out.println("<       23");
                     			System.out.println("SPACE   26");
                     			break;
         				   	case ":=":
            					System.out.println(":=      24");
                     			System.out.println("SPACE   26");
                     			break;
          				  	default: 
                                if( token.length() < 7){
                                    int j = 7 - token.length();
                                    while( j > 0){
                                        token += " ";
                                        j--;
                                    }
                                    System.out.println(token+" 28");
                                }else{
          				  		    System.out.println(token+" 28");
                                }
                     			System.out.println("SPACE   26");
          				  		break;
                     	}//end of switch
                    	 token = "";
                    }else{//just a space nothing in argument;
                    	System.out.println("SPACE   26");
       			  	}//end of if else

				}//end of for
       			if(token.equals(";") || (cArray[cArray.length - 1]) == ';'){
       			System.out.println(";       25");
       			token = "";
       		  	}
       		  	if(token.equals("END")){
       		  		System.out.println("END     31");
       		  	}
           		System.out.println("EOLN    27");

			}//end of while
		}catch(Exception e){
			e.printStackTrace();
		}// end of catch try
	}//end of main

	public static boolean isDouble(String s) {
    	try { 
    	    Double.parseDouble(s); 
    	} catch(NumberFormatException nfe) { 
    	    return false; 
    	}
    	// only got here if we didn't return false
    	return true;
	}

}