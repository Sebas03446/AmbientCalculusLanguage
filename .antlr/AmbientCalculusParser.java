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
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T__31=32, T__32=33, T__33=34, T__34=35, T__35=36, AMBIENTID=37, ID=38, 
		INT=39, STRING=40, WS=41, COMMENT=42;
	public static final int
		RULE_program = 0, RULE_statement = 1, RULE_processDeclaration = 2, RULE_processStatement = 3, 
		RULE_moveStatement = 4, RULE_ambientDeclaration = 5, RULE_conditions = 6, 
		RULE_movementStatement = 7, RULE_communicationStatement = 8, RULE_inStatement = 9, 
		RULE_outStatement = 10, RULE_openStatement = 11, RULE_callStatement = 12, 
		RULE_printStatement = 13, RULE_printConditionStatement = 14, RULE_modifyConditionStatement = 15, 
		RULE_conditionsBlock = 16, RULE_variableDeclaration = 17, RULE_conditionsVariableDeclaration = 18, 
		RULE_restriction = 19, RULE_assignmentStatement = 20, RULE_expression = 21, 
		RULE_operator = 22, RULE_comparator = 23;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statement", "processDeclaration", "processStatement", "moveStatement", 
			"ambientDeclaration", "conditions", "movementStatement", "communicationStatement", 
			"inStatement", "outStatement", "openStatement", "callStatement", "printStatement", 
			"printConditionStatement", "modifyConditionStatement", "conditionsBlock", 
			"variableDeclaration", "conditionsVariableDeclaration", "restriction", 
			"assignmentStatement", "expression", "operator", "comparator"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'process'", "'{'", "'}'", "'move'", "';'", "'ambient'", "'to'", 
			"'send'", "'receive'", "'from'", "'in'", "'out'", "'open'", "'call'", 
			"'print'", "'printc'", "'modifyc'", "'+='", "'-='", "'conditions'", "'let'", 
			"'='", "'letc'", "'restriction'", "'('", "')'", "'+'", "'-'", "'*'", 
			"'/'", "'>'", "'<'", "'>='", "'<='", "'=='", "'!='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, "AMBIENTID", "ID", "INT", "STRING", "WS", "COMMENT"
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
			setState(51);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 274880199506L) != 0)) {
				{
				{
				setState(48);
				statement();
				}
				}
				setState(53);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(54);
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
		public VariableDeclarationContext variableDeclaration() {
			return getRuleContext(VariableDeclarationContext.class,0);
		}
		public AssignmentStatementContext assignmentStatement() {
			return getRuleContext(AssignmentStatementContext.class,0);
		}
		public ModifyConditionStatementContext modifyConditionStatement() {
			return getRuleContext(ModifyConditionStatementContext.class,0);
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
			setState(68);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				enterOuterAlt(_localctx, 1);
				{
				setState(56);
				processDeclaration();
				}
				break;
			case T__5:
				enterOuterAlt(_localctx, 2);
				{
				setState(57);
				ambientDeclaration();
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 3);
				{
				setState(58);
				movementStatement();
				}
				break;
			case T__7:
			case T__8:
				enterOuterAlt(_localctx, 4);
				{
				setState(59);
				communicationStatement();
				}
				break;
			case T__10:
				enterOuterAlt(_localctx, 5);
				{
				setState(60);
				inStatement();
				}
				break;
			case T__11:
				enterOuterAlt(_localctx, 6);
				{
				setState(61);
				outStatement();
				}
				break;
			case T__12:
				enterOuterAlt(_localctx, 7);
				{
				setState(62);
				openStatement();
				}
				break;
			case T__13:
				enterOuterAlt(_localctx, 8);
				{
				setState(63);
				callStatement();
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 9);
				{
				setState(64);
				printStatement();
				}
				break;
			case T__20:
				enterOuterAlt(_localctx, 10);
				{
				setState(65);
				variableDeclaration();
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 11);
				{
				setState(66);
				assignmentStatement();
				}
				break;
			case T__16:
				enterOuterAlt(_localctx, 12);
				{
				setState(67);
				modifyConditionStatement();
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
			setState(70);
			match(T__0);
			setState(71);
			match(ID);
			setState(72);
			match(T__1);
			setState(76);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 274880265042L) != 0)) {
				{
				{
				setState(73);
				processStatement();
				}
				}
				setState(78);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(79);
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
		public PrintConditionStatementContext printConditionStatement() {
			return getRuleContext(PrintConditionStatementContext.class,0);
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
			setState(84);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(81);
				statement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(82);
				moveStatement();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(83);
				printConditionStatement();
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
			setState(86);
			match(T__3);
			setState(87);
			match(AMBIENTID);
			setState(88);
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
			setState(90);
			match(T__5);
			setState(91);
			match(AMBIENTID);
			setState(92);
			match(T__1);
			setState(94);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__19) {
				{
				setState(93);
				conditions();
				}
			}

			setState(99);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 274880199506L) != 0)) {
				{
				{
				setState(96);
				statement();
				}
				}
				setState(101);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(102);
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
		public ConditionsBlockContext conditionsBlock() {
			return getRuleContext(ConditionsBlockContext.class,0);
		}
		public List<ConditionsVariableDeclarationContext> conditionsVariableDeclaration() {
			return getRuleContexts(ConditionsVariableDeclarationContext.class);
		}
		public ConditionsVariableDeclarationContext conditionsVariableDeclaration(int i) {
			return getRuleContext(ConditionsVariableDeclarationContext.class,i);
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
			setState(104);
			conditionsBlock();
			setState(105);
			match(T__1);
			setState(109);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__22) {
				{
				{
				setState(106);
				conditionsVariableDeclaration();
				}
				}
				setState(111);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(112);
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
			setState(114);
			match(T__3);
			setState(115);
			match(ID);
			setState(116);
			match(T__6);
			setState(117);
			match(AMBIENTID);
			setState(118);
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
			setState(131);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__7:
				enterOuterAlt(_localctx, 1);
				{
				setState(120);
				match(T__7);
				setState(121);
				expression(0);
				setState(122);
				match(T__6);
				setState(123);
				match(ID);
				setState(124);
				match(T__4);
				}
				break;
			case T__8:
				enterOuterAlt(_localctx, 2);
				{
				setState(126);
				match(T__8);
				setState(127);
				match(ID);
				setState(128);
				match(T__9);
				setState(129);
				match(ID);
				setState(130);
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
			setState(133);
			match(T__10);
			setState(134);
			match(AMBIENTID);
			setState(135);
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
			setState(137);
			match(T__11);
			setState(138);
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
			setState(140);
			match(T__12);
			setState(141);
			match(AMBIENTID);
			setState(142);
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
			setState(144);
			match(T__13);
			setState(145);
			match(ID);
			setState(146);
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
			setState(148);
			match(T__14);
			setState(149);
			expression(0);
			setState(150);
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
	public static class PrintConditionStatementContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public PrintConditionStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printConditionStatement; }
	}

	public final PrintConditionStatementContext printConditionStatement() throws RecognitionException {
		PrintConditionStatementContext _localctx = new PrintConditionStatementContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_printConditionStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			match(T__15);
			setState(153);
			match(ID);
			setState(154);
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
	public static class ModifyConditionStatementContext extends ParserRuleContext {
		public Token op;
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ModifyConditionStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_modifyConditionStatement; }
	}

	public final ModifyConditionStatementContext modifyConditionStatement() throws RecognitionException {
		ModifyConditionStatementContext _localctx = new ModifyConditionStatementContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_modifyConditionStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(156);
			match(T__16);
			setState(157);
			match(ID);
			setState(158);
			((ModifyConditionStatementContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==T__17 || _la==T__18) ) {
				((ModifyConditionStatementContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(159);
			expression(0);
			setState(160);
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
	public static class ConditionsBlockContext extends ParserRuleContext {
		public ConditionsBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditionsBlock; }
	}

	public final ConditionsBlockContext conditionsBlock() throws RecognitionException {
		ConditionsBlockContext _localctx = new ConditionsBlockContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_conditionsBlock);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			match(T__19);
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
		enterRule(_localctx, 34, RULE_variableDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(164);
			match(T__20);
			setState(165);
			match(ID);
			setState(166);
			match(T__21);
			setState(167);
			expression(0);
			setState(168);
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
	public static class ConditionsVariableDeclarationContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public TerminalNode INT() { return getToken(AmbientCalculusParser.INT, 0); }
		public RestrictionContext restriction() {
			return getRuleContext(RestrictionContext.class,0);
		}
		public ConditionsVariableDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditionsVariableDeclaration; }
	}

	public final ConditionsVariableDeclarationContext conditionsVariableDeclaration() throws RecognitionException {
		ConditionsVariableDeclarationContext _localctx = new ConditionsVariableDeclarationContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_conditionsVariableDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(170);
			match(T__22);
			setState(171);
			match(ID);
			setState(172);
			match(T__21);
			setState(173);
			match(INT);
			setState(175);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__23) {
				{
				setState(174);
				restriction();
				}
			}

			setState(177);
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
	public static class RestrictionContext extends ParserRuleContext {
		public ComparatorContext comparator() {
			return getRuleContext(ComparatorContext.class,0);
		}
		public TerminalNode INT() { return getToken(AmbientCalculusParser.INT, 0); }
		public RestrictionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_restriction; }
	}

	public final RestrictionContext restriction() throws RecognitionException {
		RestrictionContext _localctx = new RestrictionContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_restriction);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(179);
			match(T__23);
			setState(180);
			comparator();
			setState(181);
			match(INT);
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
	public static class AssignmentStatementContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(AmbientCalculusParser.ID, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public AssignmentStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentStatement; }
	}

	public final AssignmentStatementContext assignmentStatement() throws RecognitionException {
		AssignmentStatementContext _localctx = new AssignmentStatementContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_assignmentStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(183);
			match(ID);
			setState(184);
			match(T__21);
			setState(185);
			expression(0);
			setState(186);
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
		int _startState = 42;
		enterRecursionRule(_localctx, 42, RULE_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(196);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__24:
				{
				setState(189);
				match(T__24);
				setState(190);
				expression(0);
				setState(191);
				match(T__25);
				}
				break;
			case ID:
				{
				setState(193);
				match(ID);
				}
				break;
			case INT:
				{
				setState(194);
				match(INT);
				}
				break;
			case STRING:
				{
				setState(195);
				match(STRING);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(204);
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
					setState(198);
					if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
					setState(199);
					operator();
					setState(200);
					expression(6);
					}
					} 
				}
				setState(206);
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
		enterRule(_localctx, 44, RULE_operator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(207);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 2013265920L) != 0)) ) {
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

	@SuppressWarnings("CheckReturnValue")
	public static class ComparatorContext extends ParserRuleContext {
		public ComparatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comparator; }
	}

	public final ComparatorContext comparator() throws RecognitionException {
		ComparatorContext _localctx = new ComparatorContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_comparator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(209);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 135291469824L) != 0)) ) {
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
		case 21:
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
		"\u0004\u0001*\u00d4\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0001\u0000\u0005\u0000"+
		"2\b\u0000\n\u0000\f\u00005\t\u0000\u0001\u0000\u0001\u0000\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001"+
		"E\b\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0005\u0002"+
		"K\b\u0002\n\u0002\f\u0002N\t\u0002\u0001\u0002\u0001\u0002\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0003\u0003U\b\u0003\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0003\u0005_\b\u0005\u0001\u0005\u0005\u0005b\b\u0005\n\u0005\f\u0005"+
		"e\t\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0005\u0006l\b\u0006\n\u0006\f\u0006o\t\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0003\b\u0084\b\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u00b0\b\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0003\u0015\u00c5\b\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0005\u0015\u00cb\b\u0015\n\u0015\f\u0015\u00ce\t\u0015\u0001"+
		"\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0000\u0001*\u0018"+
		"\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a"+
		"\u001c\u001e \"$&(*,.\u0000\u0003\u0001\u0000\u0012\u0013\u0001\u0000"+
		"\u001b\u001e\u0001\u0000\u001f$\u00d3\u00003\u0001\u0000\u0000\u0000\u0002"+
		"D\u0001\u0000\u0000\u0000\u0004F\u0001\u0000\u0000\u0000\u0006T\u0001"+
		"\u0000\u0000\u0000\bV\u0001\u0000\u0000\u0000\nZ\u0001\u0000\u0000\u0000"+
		"\fh\u0001\u0000\u0000\u0000\u000er\u0001\u0000\u0000\u0000\u0010\u0083"+
		"\u0001\u0000\u0000\u0000\u0012\u0085\u0001\u0000\u0000\u0000\u0014\u0089"+
		"\u0001\u0000\u0000\u0000\u0016\u008c\u0001\u0000\u0000\u0000\u0018\u0090"+
		"\u0001\u0000\u0000\u0000\u001a\u0094\u0001\u0000\u0000\u0000\u001c\u0098"+
		"\u0001\u0000\u0000\u0000\u001e\u009c\u0001\u0000\u0000\u0000 \u00a2\u0001"+
		"\u0000\u0000\u0000\"\u00a4\u0001\u0000\u0000\u0000$\u00aa\u0001\u0000"+
		"\u0000\u0000&\u00b3\u0001\u0000\u0000\u0000(\u00b7\u0001\u0000\u0000\u0000"+
		"*\u00c4\u0001\u0000\u0000\u0000,\u00cf\u0001\u0000\u0000\u0000.\u00d1"+
		"\u0001\u0000\u0000\u000002\u0003\u0002\u0001\u000010\u0001\u0000\u0000"+
		"\u000025\u0001\u0000\u0000\u000031\u0001\u0000\u0000\u000034\u0001\u0000"+
		"\u0000\u000046\u0001\u0000\u0000\u000053\u0001\u0000\u0000\u000067\u0005"+
		"\u0000\u0000\u00017\u0001\u0001\u0000\u0000\u00008E\u0003\u0004\u0002"+
		"\u00009E\u0003\n\u0005\u0000:E\u0003\u000e\u0007\u0000;E\u0003\u0010\b"+
		"\u0000<E\u0003\u0012\t\u0000=E\u0003\u0014\n\u0000>E\u0003\u0016\u000b"+
		"\u0000?E\u0003\u0018\f\u0000@E\u0003\u001a\r\u0000AE\u0003\"\u0011\u0000"+
		"BE\u0003(\u0014\u0000CE\u0003\u001e\u000f\u0000D8\u0001\u0000\u0000\u0000"+
		"D9\u0001\u0000\u0000\u0000D:\u0001\u0000\u0000\u0000D;\u0001\u0000\u0000"+
		"\u0000D<\u0001\u0000\u0000\u0000D=\u0001\u0000\u0000\u0000D>\u0001\u0000"+
		"\u0000\u0000D?\u0001\u0000\u0000\u0000D@\u0001\u0000\u0000\u0000DA\u0001"+
		"\u0000\u0000\u0000DB\u0001\u0000\u0000\u0000DC\u0001\u0000\u0000\u0000"+
		"E\u0003\u0001\u0000\u0000\u0000FG\u0005\u0001\u0000\u0000GH\u0005&\u0000"+
		"\u0000HL\u0005\u0002\u0000\u0000IK\u0003\u0006\u0003\u0000JI\u0001\u0000"+
		"\u0000\u0000KN\u0001\u0000\u0000\u0000LJ\u0001\u0000\u0000\u0000LM\u0001"+
		"\u0000\u0000\u0000MO\u0001\u0000\u0000\u0000NL\u0001\u0000\u0000\u0000"+
		"OP\u0005\u0003\u0000\u0000P\u0005\u0001\u0000\u0000\u0000QU\u0003\u0002"+
		"\u0001\u0000RU\u0003\b\u0004\u0000SU\u0003\u001c\u000e\u0000TQ\u0001\u0000"+
		"\u0000\u0000TR\u0001\u0000\u0000\u0000TS\u0001\u0000\u0000\u0000U\u0007"+
		"\u0001\u0000\u0000\u0000VW\u0005\u0004\u0000\u0000WX\u0005%\u0000\u0000"+
		"XY\u0005\u0005\u0000\u0000Y\t\u0001\u0000\u0000\u0000Z[\u0005\u0006\u0000"+
		"\u0000[\\\u0005%\u0000\u0000\\^\u0005\u0002\u0000\u0000]_\u0003\f\u0006"+
		"\u0000^]\u0001\u0000\u0000\u0000^_\u0001\u0000\u0000\u0000_c\u0001\u0000"+
		"\u0000\u0000`b\u0003\u0002\u0001\u0000a`\u0001\u0000\u0000\u0000be\u0001"+
		"\u0000\u0000\u0000ca\u0001\u0000\u0000\u0000cd\u0001\u0000\u0000\u0000"+
		"df\u0001\u0000\u0000\u0000ec\u0001\u0000\u0000\u0000fg\u0005\u0003\u0000"+
		"\u0000g\u000b\u0001\u0000\u0000\u0000hi\u0003 \u0010\u0000im\u0005\u0002"+
		"\u0000\u0000jl\u0003$\u0012\u0000kj\u0001\u0000\u0000\u0000lo\u0001\u0000"+
		"\u0000\u0000mk\u0001\u0000\u0000\u0000mn\u0001\u0000\u0000\u0000np\u0001"+
		"\u0000\u0000\u0000om\u0001\u0000\u0000\u0000pq\u0005\u0003\u0000\u0000"+
		"q\r\u0001\u0000\u0000\u0000rs\u0005\u0004\u0000\u0000st\u0005&\u0000\u0000"+
		"tu\u0005\u0007\u0000\u0000uv\u0005%\u0000\u0000vw\u0005\u0005\u0000\u0000"+
		"w\u000f\u0001\u0000\u0000\u0000xy\u0005\b\u0000\u0000yz\u0003*\u0015\u0000"+
		"z{\u0005\u0007\u0000\u0000{|\u0005&\u0000\u0000|}\u0005\u0005\u0000\u0000"+
		"}\u0084\u0001\u0000\u0000\u0000~\u007f\u0005\t\u0000\u0000\u007f\u0080"+
		"\u0005&\u0000\u0000\u0080\u0081\u0005\n\u0000\u0000\u0081\u0082\u0005"+
		"&\u0000\u0000\u0082\u0084\u0005\u0005\u0000\u0000\u0083x\u0001\u0000\u0000"+
		"\u0000\u0083~\u0001\u0000\u0000\u0000\u0084\u0011\u0001\u0000\u0000\u0000"+
		"\u0085\u0086\u0005\u000b\u0000\u0000\u0086\u0087\u0005%\u0000\u0000\u0087"+
		"\u0088\u0005\u0005\u0000\u0000\u0088\u0013\u0001\u0000\u0000\u0000\u0089"+
		"\u008a\u0005\f\u0000\u0000\u008a\u008b\u0005\u0005\u0000\u0000\u008b\u0015"+
		"\u0001\u0000\u0000\u0000\u008c\u008d\u0005\r\u0000\u0000\u008d\u008e\u0005"+
		"%\u0000\u0000\u008e\u008f\u0005\u0005\u0000\u0000\u008f\u0017\u0001\u0000"+
		"\u0000\u0000\u0090\u0091\u0005\u000e\u0000\u0000\u0091\u0092\u0005&\u0000"+
		"\u0000\u0092\u0093\u0005\u0005\u0000\u0000\u0093\u0019\u0001\u0000\u0000"+
		"\u0000\u0094\u0095\u0005\u000f\u0000\u0000\u0095\u0096\u0003*\u0015\u0000"+
		"\u0096\u0097\u0005\u0005\u0000\u0000\u0097\u001b\u0001\u0000\u0000\u0000"+
		"\u0098\u0099\u0005\u0010\u0000\u0000\u0099\u009a\u0005&\u0000\u0000\u009a"+
		"\u009b\u0005\u0005\u0000\u0000\u009b\u001d\u0001\u0000\u0000\u0000\u009c"+
		"\u009d\u0005\u0011\u0000\u0000\u009d\u009e\u0005&\u0000\u0000\u009e\u009f"+
		"\u0007\u0000\u0000\u0000\u009f\u00a0\u0003*\u0015\u0000\u00a0\u00a1\u0005"+
		"\u0005\u0000\u0000\u00a1\u001f\u0001\u0000\u0000\u0000\u00a2\u00a3\u0005"+
		"\u0014\u0000\u0000\u00a3!\u0001\u0000\u0000\u0000\u00a4\u00a5\u0005\u0015"+
		"\u0000\u0000\u00a5\u00a6\u0005&\u0000\u0000\u00a6\u00a7\u0005\u0016\u0000"+
		"\u0000\u00a7\u00a8\u0003*\u0015\u0000\u00a8\u00a9\u0005\u0005\u0000\u0000"+
		"\u00a9#\u0001\u0000\u0000\u0000\u00aa\u00ab\u0005\u0017\u0000\u0000\u00ab"+
		"\u00ac\u0005&\u0000\u0000\u00ac\u00ad\u0005\u0016\u0000\u0000\u00ad\u00af"+
		"\u0005\'\u0000\u0000\u00ae\u00b0\u0003&\u0013\u0000\u00af\u00ae\u0001"+
		"\u0000\u0000\u0000\u00af\u00b0\u0001\u0000\u0000\u0000\u00b0\u00b1\u0001"+
		"\u0000\u0000\u0000\u00b1\u00b2\u0005\u0005\u0000\u0000\u00b2%\u0001\u0000"+
		"\u0000\u0000\u00b3\u00b4\u0005\u0018\u0000\u0000\u00b4\u00b5\u0003.\u0017"+
		"\u0000\u00b5\u00b6\u0005\'\u0000\u0000\u00b6\'\u0001\u0000\u0000\u0000"+
		"\u00b7\u00b8\u0005&\u0000\u0000\u00b8\u00b9\u0005\u0016\u0000\u0000\u00b9"+
		"\u00ba\u0003*\u0015\u0000\u00ba\u00bb\u0005\u0005\u0000\u0000\u00bb)\u0001"+
		"\u0000\u0000\u0000\u00bc\u00bd\u0006\u0015\uffff\uffff\u0000\u00bd\u00be"+
		"\u0005\u0019\u0000\u0000\u00be\u00bf\u0003*\u0015\u0000\u00bf\u00c0\u0005"+
		"\u001a\u0000\u0000\u00c0\u00c5\u0001\u0000\u0000\u0000\u00c1\u00c5\u0005"+
		"&\u0000\u0000\u00c2\u00c5\u0005\'\u0000\u0000\u00c3\u00c5\u0005(\u0000"+
		"\u0000\u00c4\u00bc\u0001\u0000\u0000\u0000\u00c4\u00c1\u0001\u0000\u0000"+
		"\u0000\u00c4\u00c2\u0001\u0000\u0000\u0000\u00c4\u00c3\u0001\u0000\u0000"+
		"\u0000\u00c5\u00cc\u0001\u0000\u0000\u0000\u00c6\u00c7\n\u0005\u0000\u0000"+
		"\u00c7\u00c8\u0003,\u0016\u0000\u00c8\u00c9\u0003*\u0015\u0006\u00c9\u00cb"+
		"\u0001\u0000\u0000\u0000\u00ca\u00c6\u0001\u0000\u0000\u0000\u00cb\u00ce"+
		"\u0001\u0000\u0000\u0000\u00cc\u00ca\u0001\u0000\u0000\u0000\u00cc\u00cd"+
		"\u0001\u0000\u0000\u0000\u00cd+\u0001\u0000\u0000\u0000\u00ce\u00cc\u0001"+
		"\u0000\u0000\u0000\u00cf\u00d0\u0007\u0001\u0000\u0000\u00d0-\u0001\u0000"+
		"\u0000\u0000\u00d1\u00d2\u0007\u0002\u0000\u0000\u00d2/\u0001\u0000\u0000"+
		"\u0000\u000b3DLT^cm\u0083\u00af\u00c4\u00cc";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}