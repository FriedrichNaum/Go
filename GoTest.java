import java.io.FileInputStream;
import java.io.FileNotFoundException;

public class GoTest {
    public static void main(String[] args) throws ParseException, FileNotFoundException {
        Go lexer = new Go(new FileInputStream("input.go"));
        Token token;
        while ((token = lexer.getNextToken()).kind != GoConstants.EOF) {
            switch (token.kind) {
                case GoConstants.ID:
                    System.out.println("ID " + token.image);
                    break;
                case GoConstants.NUM_DEC:
                    System.out.println("NUM_DEC " + token.image);
                    break;
                case GoConstants.KEYWORD:
                    System.out.println("KEYWORD " + token.image);
                    break;
                case GoConstants.OPERATOR:
                    System.out.println("OPERATOR " + token.image);
                    break;
                case GoConstants.DELIMITER:
                    System.out.println("DELIMITER " + token.image);
                    break;
                case GoConstants.STRING:
                    System.out.println("STRING " + token.image);
                    break;
                case GoConstants.SINGLE_LINE_COMMENT:
                    System.out.println("SINGLE_LINE_COMMENT " + token.image);
                    break;
                case GoConstants.MULTI_LINE_COMMENT:
                    System.out.println("MULTI_LINE_COMMENT " + token.image);
                    break;
                    case GoConstants.WHITESPACE:
                    // Ignore whitespace tokens
                    break;
                default:
                    System.out.println("UNKNOWN " + token.image);
                    break;
            }
        }
    }
}
