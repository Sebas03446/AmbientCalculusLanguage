// Generated from /home/jeansebastiansoracipa/GitHub/UNAL/AmbientCalculusLanguage/AmbientCalculus.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link AmbientCalculusParser}.
 */
public interface AmbientCalculusListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link AmbientCalculusParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(AmbientCalculusParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link AmbientCalculusParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(AmbientCalculusParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link AmbientCalculusParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(AmbientCalculusParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link AmbientCalculusParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(AmbientCalculusParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link AmbientCalculusParser#processDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterProcessDeclaration(AmbientCalculusParser.ProcessDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link AmbientCalculusParser#processDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitProcessDeclaration(AmbientCalculusParser.ProcessDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link AmbientCalculusParser#processStatement}.
	 * @param ctx the parse tree
	 */
	void enterProcessStatement(AmbientCalculusParser.ProcessStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link AmbientCalculusParser#processStatement}.
	 * @param ctx the parse tree
	 */
	void exitProcessStatement(AmbientCalculusParser.ProcessStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link AmbientCalculusParser#moveStatement}.
	 * @param ctx the parse tree
	 */
	void enterMoveStatement(AmbientCalculusParser.MoveStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link AmbientCalculusParser#moveStatement}.
	 * @param ctx the parse tree
	 */
	void exitMoveStatement(AmbientCalculusParser.MoveStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link AmbientCalculusParser#ambientDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterAmbientDeclaration(AmbientCalculusParser.AmbientDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link AmbientCalculusParser#ambientDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitAmbientDeclaration(AmbientCalculusParser.AmbientDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link AmbientCalculusParser#conditions}.
	 * @param ctx the parse tree
	 */
	void enterConditions(AmbientCalculusParser.ConditionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link AmbientCalculusParser#conditions}.
	 * @param ctx the parse tree
	 */
	void exitConditions(AmbientCalculusParser.ConditionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link AmbientCalculusParser#movementStatement}.
	 * @param ctx the parse tree
	 */
	void enterMovementStatement(AmbientCalculusParser.MovementStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link AmbientCalculusParser#movementStatement}.
	 * @param ctx the parse tree
	 */
	void exitMovementStatement(AmbientCalculusParser.MovementStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link AmbientCalculusParser#communicationStatement}.
	 * @param ctx the parse tree
	 */
	void enterCommunicationStatement(AmbientCalculusParser.CommunicationStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link AmbientCalculusParser#communicationStatement}.
	 * @param ctx the parse tree
	 */
	void exitCommunicationStatement(AmbientCalculusParser.CommunicationStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link AmbientCalculusParser#inStatement}.
	 * @param ctx the parse tree
	 */
	void enterInStatement(AmbientCalculusParser.InStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link AmbientCalculusParser#inStatement}.
	 * @param ctx the parse tree
	 */
	void exitInStatement(AmbientCalculusParser.InStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link AmbientCalculusParser#outStatement}.
	 * @param ctx the parse tree
	 */
	void enterOutStatement(AmbientCalculusParser.OutStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link AmbientCalculusParser#outStatement}.
	 * @param ctx the parse tree
	 */
	void exitOutStatement(AmbientCalculusParser.OutStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link AmbientCalculusParser#openStatement}.
	 * @param ctx the parse tree
	 */
	void enterOpenStatement(AmbientCalculusParser.OpenStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link AmbientCalculusParser#openStatement}.
	 * @param ctx the parse tree
	 */
	void exitOpenStatement(AmbientCalculusParser.OpenStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link AmbientCalculusParser#callStatement}.
	 * @param ctx the parse tree
	 */
	void enterCallStatement(AmbientCalculusParser.CallStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link AmbientCalculusParser#callStatement}.
	 * @param ctx the parse tree
	 */
	void exitCallStatement(AmbientCalculusParser.CallStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link AmbientCalculusParser#printStatement}.
	 * @param ctx the parse tree
	 */
	void enterPrintStatement(AmbientCalculusParser.PrintStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link AmbientCalculusParser#printStatement}.
	 * @param ctx the parse tree
	 */
	void exitPrintStatement(AmbientCalculusParser.PrintStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link AmbientCalculusParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfStatement(AmbientCalculusParser.IfStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link AmbientCalculusParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfStatement(AmbientCalculusParser.IfStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link AmbientCalculusParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterVariableDeclaration(AmbientCalculusParser.VariableDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link AmbientCalculusParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitVariableDeclaration(AmbientCalculusParser.VariableDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link AmbientCalculusParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(AmbientCalculusParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link AmbientCalculusParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(AmbientCalculusParser.ExpressionContext ctx);
}