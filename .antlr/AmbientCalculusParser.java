// Generated from /home/jeansebastiansoracipa/GitHub/UNAL/AmbientCalculusLanguage/AmbientCalculus.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class AmbientCalculusParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, AMBIENTID=26, ID=27, INT=28, STRING=29, WS=30, COMMENT=31;
	public static final int
		RULE_program = 0, RULE_statement = 1, RULE_processDeclaration = 2, RULE_processStatement = 3, 
		RULE_moveStatement = 4, RULE_ambientDeclaration = 5, RULE_conditions = 6, 
		RULE_movementStatement = 7, RULE_communicationStatement = 8, RULE_inStatement = 9, 
		RULE_outStatement = 10, RULE_openStatement = 11, RULE_callStatement = 12, 
		RULE_printStatement = 13, RULE_ifStatement = 14, RULE_variableDeclaration = 15, 
		RULE_expression = 16, RULE_operator = 17;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statement", "processDeclaration", "processStatement", "moveStatement", 
			"ambientDeclaration", "conditions", "movementStatement", "communicationStatement", 
			"inStatement", "outStatement", "openStatement", "callStatement", "printStatement", 
			"ifStatement", "variableDeclaration", "expression", "operator"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'process'", "'{'", "'}'", "'move'", "';'", "'ambient'", "'conditions'", 
			"'to'", "'send'", "'receive'", "'from'", "'in'", "'out'", "'open'", "'call'", 
			"'print'", "'if'", "'('", "')'", "'let'", "'='", "'+'", "'-'", "'*'", 
			"'/'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, "AMBIENTID", "ID", "INT", "STRING", "WS", "COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "AmbientCalculus.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public AmbientCalculusParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(AmbientCalculusParser.EOF, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(39);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1308242L) != 0)) {
				{
				{
				setState(36);
				statement();
				}
				}
				setState(41);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(42);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public ProcessDeclarationContext processDeclaration() {
			return getRuleContext(ProcessDeclarationContext.class,0);
		}
		public AmbientDeclarationContext ambientDeclaration() {
			return getRuleContext(AmbientDeclarationContext.class,0);
		}
		public MovementStatementContext movementStatement() {
			return getRuleContext(MovementStatementContext.class,0);
		}
		public CommunicationStatementContext communicationStatement() {
			return getRuleContext(CommunicationStatementContext.class,0);
		}
		public InStatementContext inStatement() {
			return getRuleContext(InStatementContext.class,0);
		}
		public OutStatementContext outStatement() {
			return getRuleContext(OutStatementContext.class,0);
		}
		public OpenStatementContext openStatement() {
			return getRuleContext(OpenStatementContext.class,0);
		}
		public CallStatementContext callStatement() {
			return getRuleContext(CallStatementContext.class,0);
		}
		public PrintStatementContext printStatement() {
			return getRuleContext(PrintStatementContext.class,0);
		}
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public VariableDeclarationContext variableDeclaration() {
			return getRuleContext(VariableDeclarationContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statement);
		try {
			setState(55);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				enterOuterAlt(_localctx, 1);
				{
				setState(44);
				processDeclaration();
				}
				break;
			case T__5:
				enterOuterAlt(_localctx, 2);
				{
				setState(45);
				ambientDeclaration();
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 3);
				{
				setState(46);
				movementStatement();
				}
				break;
			case T__8:
			case T__9:
				enterOuterAlt(_localctx, 4);
				{
				setState(47);
				communicationStatement();
				}
				break;
			case T__11:
				enterOuterAlt(_localctx, 5);
				{
				setState(48);
				inStatement();
				}
				break;
			case T__12:
				enterOuterAlt(_localctx, 6);
				{
				setState(49);
				outStatement();
				}
				break;
			case T__13:
				enterOuterAlt(_localctx, 7);
				{
				setState(50);
				openStatement();
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 8);
				{
				setState(51);
				callStatement();
				}
				break;
			case T__15:
				enterOuterAlt(_localctx, 9);
				{
				setState(52);
				printStatement();
				}
				break;
			case T__16:
				enterOuterAlt(_localctx, 10);
				{
				setState(53);
				ifStatement();
				}
				break;
			case T__19:
				enterOuterAlt(_localctx, 11);
				{
				setState(54);
				variableDeclaration();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProcessDeclarationContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public List<ProcessStatementContext> processStatement() {
			return getRuleContexts(ProcessStatementContext.class);
		}
		public ProcessStatementContext processStatement(int i) {
			return getRuleContext(ProcessStatementContext.class,i);
		}
		public ProcessDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_processDeclaration; }
	}

	public final ProcessDeclarationContext processDeclaration() throws RecognitionException {
		ProcessDeclarationContext _localctx = new ProcessDeclarationContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_processDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(57);
			match(T__0);
			setState(58);
			match(ID);
			setState(59);
			match(T__1);
			setState(63);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1308242L) != 0)) {
				{
				{
				setState(60);
				processStatement();
				}
				}
				setState(65);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(66);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProcessStatementContext extends ParserRuleContext {
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public MoveStatementContext moveStatement() {
			return getRuleContext(MoveStatementContext.class,0);
		}
		public ProcessStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_processStatement; }
	}

	public final ProcessStatementContext processStatement() throws RecognitionException {
		ProcessStatementContext _localctx = new ProcessStatementContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_processStatement);
		try {
			setState(70);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(68);
				statement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(69);
				moveStatement();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MoveStatementContext extends ParserRuleContext {
		public TerminalNode AMBIENTID() { return getToken(AmbientCalculusParser.AMBIENTID, 0); }
		public MoveStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_moveStatement; }
	}

	public final MoveStatementContext moveStatement() throws RecognitionException {
		MoveStatementContext _localctx = new MoveStatementContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_moveStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(72);
			match(T__3);
			setState(73);
			match(AMBIENTID);
			setState(74);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AmbientDeclarationContext extends ParserRuleContext {
		public TerminalNode AMBIENTID() { return getToken(AmbientCalculusParser.AMBIENTID, 0); }
		public ConditionsContext conditions() {
			return getRuleContext(ConditionsContext.class,0);
		}
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public AmbientDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ambientDeclaration; }
	}

	public final AmbientDeclarationContext ambientDeclaration() throws RecognitionException {
		AmbientDeclarationContext _localctx = new AmbientDeclarationContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_ambientDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(76);
			match(T__5);
			setState(77);
			match(AMBIENTID);
			setState(78);
			match(T__1);
			setState(80);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(79);
				conditions();
				}
			}

			setState(85);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1308242L) != 0)) {
				{
				{
				setState(82);
				statement();
				}
				}
				setState(87);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(88);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ConditionsContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public ConditionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditions; }
	}

	public final ConditionsContext conditions() throws RecognitionException {
		ConditionsContext _localctx = new ConditionsContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_conditions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(90);
			match(T__6);
			setState(91);
			match(T__1);
			setState(95);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1308242L) != 0)) {
				{
				{
				setState(92);
				statement();
				}
				}
				setState(97);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(98);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MovementStatementContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public TerminalNode AMBIENTID() { return getToken(AmbientCalculusParser.AMBIENTID, 0); }
		public MovementStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_movementStatement; }
	}

	public final MovementStatementContext movementStatement() throws RecognitionException {
		MovementStatementContext _localctx = new MovementStatementContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_movementStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(100);
			match(T__3);
			setState(101);
			match(ID);
			setState(102);
			match(T__7);
			setState(103);
			match(AMBIENTID);
			setState(104);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CommunicationStatementContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<TerminalNode> ID() { return getTokens(AmbientCalculusParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(AmbientCalculusParser.ID, i);
		}
		public CommunicationStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_communicationStatement; }
	}

	public final CommunicationStatementContext communicationStatement() throws RecognitionException {
		CommunicationStatementContext _localctx = new CommunicationStatementContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_communicationStatement);
		try {
			setState(117);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__8:
				enterOuterAlt(_localctx, 1);
				{
				setState(106);
				match(T__8);
				setState(107);
				expression(0);
				setState(108);
				match(T__7);
				setState(109);
				match(ID);
				setState(110);
				match(T__4);
				}
				break;
			case T__9:
				enterOuterAlt(_localctx, 2);
				{
				setState(112);
				match(T__9);
				setState(113);
				match(ID);
				setState(114);
				match(T__10);
				setState(115);
				match(ID);
				setState(116);
				match(T__4);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InStatementContext extends ParserRuleContext {
		public TerminalNode AMBIENTID() { return getToken(AmbientCalculusParser.AMBIENTID, 0); }
		public InStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_inStatement; }
	}

	public final InStatementContext inStatement() throws RecognitionException {
		InStatementContext _localctx = new InStatementContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_inStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(119);
			match(T__11);
			setState(120);
			match(AMBIENTID);
			setState(121);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OutStatementContext extends ParserRuleContext {
		public OutStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_outStatement; }
	}

	public final OutStatementContext outStatement() throws RecognitionException {
		OutStatementContext _localctx = new OutStatementContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_outStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(123);
			match(T__12);
			setState(124);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OpenStatementContext extends ParserRuleContext {
		public TerminalNode AMBIENTID() { return getToken(AmbientCalculusParser.AMBIENTID, 0); }
		public OpenStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_openStatement; }
	}

	public final OpenStatementContext openStatement() throws RecognitionException {
		OpenStatementContext _localctx = new OpenStatementContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_openStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(126);
			match(T__13);
			setState(127);
			match(AMBIENTID);
			setState(128);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CallStatementContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public CallStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_callStatement; }
	}

	public final CallStatementContext callStatement() throws RecognitionException {
		CallStatementContext _localctx = new CallStatementContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_callStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(130);
			match(T__14);
			setState(131);
			match(ID);
			setState(132);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrintStatementContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public PrintStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printStatement; }
	}

	public final PrintStatementContext printStatement() throws RecognitionException {
		PrintStatementContext _localctx = new PrintStatementContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_printStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
			match(T__15);
			setState(135);
			expression(0);
			setState(136);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IfStatementContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public IfStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStatement; }
	}

	public final IfStatementContext ifStatement() throws RecognitionException {
		IfStatementContext _localctx = new IfStatementContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_ifStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(138);
			match(T__16);
			setState(139);
			match(T__17);
			setState(140);
			expression(0);
			setState(141);
			match(T__18);
			setState(142);
			match(T__1);
			setState(146);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1308242L) != 0)) {
				{
				{
				setState(143);
				statement();
				}
				}
				setState(148);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(149);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VariableDeclarationContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public VariableDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variableDeclaration; }
	}

	public final VariableDeclarationContext variableDeclaration() throws RecognitionException {
		VariableDeclarationContext _localctx = new VariableDeclarationContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_variableDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(151);
			match(T__19);
			setState(152);
			match(ID);
			setState(153);
			match(T__20);
			setState(154);
			expression(0);
			setState(155);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public TerminalNode INT() { return getToken(AmbientCalculusParser.INT, 0); }
		public TerminalNode STRING() { return getToken(AmbientCalculusParser.STRING, 0); }
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 32;
		enterRecursionRule(_localctx, 32, RULE_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(165);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__17:
				{
				setState(158);
				match(T__17);
				setState(159);
				expression(0);
				setState(160);
				match(T__18);
				}
				break;
			case ID:
				{
				setState(162);
				match(ID);
				}
				break;
			case INT:
				{
				setState(163);
				match(INT);
				}
				break;
			case STRING:
				{
				setState(164);
				match(STRING);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(173);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ExpressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_expression);
					setState(167);
					if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
					setState(168);
					operator();
					setState(169);
					expression(6);
					}
					} 
				}
				setState(175);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OperatorContext extends ParserRuleContext {
		public OperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator; }
	}

	public final OperatorContext operator() throws RecognitionException {
		OperatorContext _localctx = new OperatorContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_operator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(176);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 62914560L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 16:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 5);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u001f\u00b3\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001"+
		"\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004"+
		"\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007"+
		"\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b"+
		"\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007"+
		"\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0001\u0000\u0005"+
		"\u0000&\b\u0000\n\u0000\f\u0000)\t\u0000\u0001\u0000\u0001\u0000\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u00018\b"+
		"\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0005\u0002>\b"+
		"\u0002\n\u0002\f\u0002A\t\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001"+
		"\u0003\u0003\u0003G\b\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005Q\b"+
		"\u0005\u0001\u0005\u0005\u0005T\b\u0005\n\u0005\f\u0005W\t\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0005\u0006^\b"+
		"\u0006\n\u0006\f\u0006a\t\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b"+
		"\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0003\bv\b\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001"+
		"\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0005\u000e\u0091\b\u000e\n"+
		"\u000e\f\u000e\u0094\t\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0003\u0010\u00a6\b\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0005\u0010\u00ac\b\u0010\n\u0010\f\u0010\u00af\t\u0010\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0000\u0001 \u0012\u0000\u0002\u0004\u0006\b"+
		"\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"\u0000\u0001"+
		"\u0001\u0000\u0016\u0019\u00b6\u0000\'\u0001\u0000\u0000\u0000\u00027"+
		"\u0001\u0000\u0000\u0000\u00049\u0001\u0000\u0000\u0000\u0006F\u0001\u0000"+
		"\u0000\u0000\bH\u0001\u0000\u0000\u0000\nL\u0001\u0000\u0000\u0000\fZ"+
		"\u0001\u0000\u0000\u0000\u000ed\u0001\u0000\u0000\u0000\u0010u\u0001\u0000"+
		"\u0000\u0000\u0012w\u0001\u0000\u0000\u0000\u0014{\u0001\u0000\u0000\u0000"+
		"\u0016~\u0001\u0000\u0000\u0000\u0018\u0082\u0001\u0000\u0000\u0000\u001a"+
		"\u0086\u0001\u0000\u0000\u0000\u001c\u008a\u0001\u0000\u0000\u0000\u001e"+
		"\u0097\u0001\u0000\u0000\u0000 \u00a5\u0001\u0000\u0000\u0000\"\u00b0"+
		"\u0001\u0000\u0000\u0000$&\u0003\u0002\u0001\u0000%$\u0001\u0000\u0000"+
		"\u0000&)\u0001\u0000\u0000\u0000\'%\u0001\u0000\u0000\u0000\'(\u0001\u0000"+
		"\u0000\u0000(*\u0001\u0000\u0000\u0000)\'\u0001\u0000\u0000\u0000*+\u0005"+
		"\u0000\u0000\u0001+\u0001\u0001\u0000\u0000\u0000,8\u0003\u0004\u0002"+
		"\u0000-8\u0003\n\u0005\u0000.8\u0003\u000e\u0007\u0000/8\u0003\u0010\b"+
		"\u000008\u0003\u0012\t\u000018\u0003\u0014\n\u000028\u0003\u0016\u000b"+
		"\u000038\u0003\u0018\f\u000048\u0003\u001a\r\u000058\u0003\u001c\u000e"+
		"\u000068\u0003\u001e\u000f\u00007,\u0001\u0000\u0000\u00007-\u0001\u0000"+
		"\u0000\u00007.\u0001\u0000\u0000\u00007/\u0001\u0000\u0000\u000070\u0001"+
		"\u0000\u0000\u000071\u0001\u0000\u0000\u000072\u0001\u0000\u0000\u0000"+
		"73\u0001\u0000\u0000\u000074\u0001\u0000\u0000\u000075\u0001\u0000\u0000"+
		"\u000076\u0001\u0000\u0000\u00008\u0003\u0001\u0000\u0000\u00009:\u0005"+
		"\u0001\u0000\u0000:;\u0005\u001b\u0000\u0000;?\u0005\u0002\u0000\u0000"+
		"<>\u0003\u0006\u0003\u0000=<\u0001\u0000\u0000\u0000>A\u0001\u0000\u0000"+
		"\u0000?=\u0001\u0000\u0000\u0000?@\u0001\u0000\u0000\u0000@B\u0001\u0000"+
		"\u0000\u0000A?\u0001\u0000\u0000\u0000BC\u0005\u0003\u0000\u0000C\u0005"+
		"\u0001\u0000\u0000\u0000DG\u0003\u0002\u0001\u0000EG\u0003\b\u0004\u0000"+
		"FD\u0001\u0000\u0000\u0000FE\u0001\u0000\u0000\u0000G\u0007\u0001\u0000"+
		"\u0000\u0000HI\u0005\u0004\u0000\u0000IJ\u0005\u001a\u0000\u0000JK\u0005"+
		"\u0005\u0000\u0000K\t\u0001\u0000\u0000\u0000LM\u0005\u0006\u0000\u0000"+
		"MN\u0005\u001a\u0000\u0000NP\u0005\u0002\u0000\u0000OQ\u0003\f\u0006\u0000"+
		"PO\u0001\u0000\u0000\u0000PQ\u0001\u0000\u0000\u0000QU\u0001\u0000\u0000"+
		"\u0000RT\u0003\u0002\u0001\u0000SR\u0001\u0000\u0000\u0000TW\u0001\u0000"+
		"\u0000\u0000US\u0001\u0000\u0000\u0000UV\u0001\u0000\u0000\u0000VX\u0001"+
		"\u0000\u0000\u0000WU\u0001\u0000\u0000\u0000XY\u0005\u0003\u0000\u0000"+
		"Y\u000b\u0001\u0000\u0000\u0000Z[\u0005\u0007\u0000\u0000[_\u0005\u0002"+
		"\u0000\u0000\\^\u0003\u0002\u0001\u0000]\\\u0001\u0000\u0000\u0000^a\u0001"+
		"\u0000\u0000\u0000_]\u0001\u0000\u0000\u0000_`\u0001\u0000\u0000\u0000"+
		"`b\u0001\u0000\u0000\u0000a_\u0001\u0000\u0000\u0000bc\u0005\u0003\u0000"+
		"\u0000c\r\u0001\u0000\u0000\u0000de\u0005\u0004\u0000\u0000ef\u0005\u001b"+
		"\u0000\u0000fg\u0005\b\u0000\u0000gh\u0005\u001a\u0000\u0000hi\u0005\u0005"+
		"\u0000\u0000i\u000f\u0001\u0000\u0000\u0000jk\u0005\t\u0000\u0000kl\u0003"+
		" \u0010\u0000lm\u0005\b\u0000\u0000mn\u0005\u001b\u0000\u0000no\u0005"+
		"\u0005\u0000\u0000ov\u0001\u0000\u0000\u0000pq\u0005\n\u0000\u0000qr\u0005"+
		"\u001b\u0000\u0000rs\u0005\u000b\u0000\u0000st\u0005\u001b\u0000\u0000"+
		"tv\u0005\u0005\u0000\u0000uj\u0001\u0000\u0000\u0000up\u0001\u0000\u0000"+
		"\u0000v\u0011\u0001\u0000\u0000\u0000wx\u0005\f\u0000\u0000xy\u0005\u001a"+
		"\u0000\u0000yz\u0005\u0005\u0000\u0000z\u0013\u0001\u0000\u0000\u0000"+
		"{|\u0005\r\u0000\u0000|}\u0005\u0005\u0000\u0000}\u0015\u0001\u0000\u0000"+
		"\u0000~\u007f\u0005\u000e\u0000\u0000\u007f\u0080\u0005\u001a\u0000\u0000"+
		"\u0080\u0081\u0005\u0005\u0000\u0000\u0081\u0017\u0001\u0000\u0000\u0000"+
		"\u0082\u0083\u0005\u000f\u0000\u0000\u0083\u0084\u0005\u001b\u0000\u0000"+
		"\u0084\u0085\u0005\u0005\u0000\u0000\u0085\u0019\u0001\u0000\u0000\u0000"+
		"\u0086\u0087\u0005\u0010\u0000\u0000\u0087\u0088\u0003 \u0010\u0000\u0088"+
		"\u0089\u0005\u0005\u0000\u0000\u0089\u001b\u0001\u0000\u0000\u0000\u008a"+
		"\u008b\u0005\u0011\u0000\u0000\u008b\u008c\u0005\u0012\u0000\u0000\u008c"+
		"\u008d\u0003 \u0010\u0000\u008d\u008e\u0005\u0013\u0000\u0000\u008e\u0092"+
		"\u0005\u0002\u0000\u0000\u008f\u0091\u0003\u0002\u0001\u0000\u0090\u008f"+
		"\u0001\u0000\u0000\u0000\u0091\u0094\u0001\u0000\u0000\u0000\u0092\u0090"+
		"\u0001\u0000\u0000\u0000\u0092\u0093\u0001\u0000\u0000\u0000\u0093\u0095"+
		"\u0001\u0000\u0000\u0000\u0094\u0092\u0001\u0000\u0000\u0000\u0095\u0096"+
		"\u0005\u0003\u0000\u0000\u0096\u001d\u0001\u0000\u0000\u0000\u0097\u0098"+
		"\u0005\u0014\u0000\u0000\u0098\u0099\u0005\u001b\u0000\u0000\u0099\u009a"+
		"\u0005\u0015\u0000\u0000\u009a\u009b\u0003 \u0010\u0000\u009b\u009c\u0005"+
		"\u0005\u0000\u0000\u009c\u001f\u0001\u0000\u0000\u0000\u009d\u009e\u0006"+
		"\u0010\uffff\uffff\u0000\u009e\u009f\u0005\u0012\u0000\u0000\u009f\u00a0"+
		"\u0003 \u0010\u0000\u00a0\u00a1\u0005\u0013\u0000\u0000\u00a1\u00a6\u0001"+
		"\u0000\u0000\u0000\u00a2\u00a6\u0005\u001b\u0000\u0000\u00a3\u00a6\u0005"+
		"\u001c\u0000\u0000\u00a4\u00a6\u0005\u001d\u0000\u0000\u00a5\u009d\u0001"+
		"\u0000\u0000\u0000\u00a5\u00a2\u0001\u0000\u0000\u0000\u00a5\u00a3\u0001"+
		"\u0000\u0000\u0000\u00a5\u00a4\u0001\u0000\u0000\u0000\u00a6\u00ad\u0001"+
		"\u0000\u0000\u0000\u00a7\u00a8\n\u0005\u0000\u0000\u00a8\u00a9\u0003\""+
		"\u0011\u0000\u00a9\u00aa\u0003 \u0010\u0006\u00aa\u00ac\u0001\u0000\u0000"+
		"\u0000\u00ab\u00a7\u0001\u0000\u0000\u0000\u00ac\u00af\u0001\u0000\u0000"+
		"\u0000\u00ad\u00ab\u0001\u0000\u0000\u0000\u00ad\u00ae\u0001\u0000\u0000"+
		"\u0000\u00ae!\u0001\u0000\u0000\u0000\u00af\u00ad\u0001\u0000\u0000\u0000"+
		"\u00b0\u00b1\u0007\u0000\u0000\u0000\u00b1#\u0001\u0000\u0000\u0000\u000b"+
		"\'7?FPU_u\u0092\u00a5\u00ad";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}